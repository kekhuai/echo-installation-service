package handler

import (
	"net/http"
	"strconv"

	"github.com/kekhuay/echo-installation-service/model"
	"github.com/kekhuay/echo-installation-service/utils"
	"github.com/labstack/echo"
)

// CreatePartner create new partner
func (h *Handler) CreatePartner(c echo.Context) error {
	var p model.Partner
	req := &partnerCreateRequest{}
	if err := req.bind(c, &p); nil != err {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := h.partnerStore.CreatePartner(&p)
	if nil != err {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newPartnerResponse(c, &p))
}

// Partners get all partners
func (h *Handler) Partners(c echo.Context) error {
	offset, err := strconv.ParseInt(c.QueryParam("offset"), 10, 64)
	if nil != err {
		offset = 0
	}
	limit, err := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if nil != err {
		limit = 20
	}
	var partners []model.Partner
	var count int64
	partners, count, err = h.partnerStore.List(offset, limit)
	if nil != err {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newPartnerListResponse(partners, count))
}

// UpdatePartner update a partner by id
func (h *Handler) UpdatePartner(c echo.Context) error {
	id := c.Param("id")
	partner, err := h.partnerStore.GetByID(id)
	if nil != err {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if nil == partner {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &partnerUpdateRequest{}
	req.populate(partner)
	if err := req.bind(c, partner); nil != err {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.partnerStore.UpdatePartner(partner); nil != err {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newPartnerResponse(c, partner))
}

// DeletePartner delete a partner by id
func (h *Handler) DeletePartner(c echo.Context) error {
	id := c.Param("id")
	partner, err := h.partnerStore.GetByID(id)
	if nil != err {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if nil == partner {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.partnerStore.DeletePartner(partner)
	if nil != err {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
