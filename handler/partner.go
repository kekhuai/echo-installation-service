package handler

import (
	"net/http"

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
