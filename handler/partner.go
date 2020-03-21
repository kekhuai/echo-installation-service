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
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, newPartnerListResponse(partners, count))
}
