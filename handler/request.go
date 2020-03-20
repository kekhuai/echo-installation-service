package handler

import (
	"github.com/kekhuay/echo-installation-service/model"
	"github.com/labstack/echo"
)

type partnerCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r *partnerCreateRequest) bind(c echo.Context, p *model.Partner) error {
	if err := c.Bind(r); nil != err {
		return err
	}
	if err := c.Validate(r); nil != err {
		return err
	}
	p.Name = r.Name
	return nil
}
