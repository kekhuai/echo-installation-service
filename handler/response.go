package handler

import (
	"github.com/kekhuay/echo-installation-service/model"
	"github.com/labstack/echo"
)

type partnerResponse struct {
	Name string `json:"name"`
}

func newPartnerResponse(c echo.Context, p *model.Partner) *partnerResponse {
	pr := new(partnerResponse)
	pr.Name = p.Name
	return pr
}
