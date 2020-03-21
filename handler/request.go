package handler

import (
	"github.com/kekhuay/echo-installation-service/model"
	"github.com/labstack/echo"
)

type partnerCreateRequest struct {
	Partner struct {
		Name string `json:"name"`
	} `json:"partner"`
}

type partnerUpdateRequest struct {
	Partner struct {
		Name string `json:"name"`
	} `json:"partner"`
}

func (r *partnerCreateRequest) bind(c echo.Context, p *model.Partner) error {
	if err := c.Bind(r); nil != err {
		return err
	}
	if err := c.Validate(r); nil != err {
		return err
	}
	p.Name = r.Partner.Name
	return nil
}

func (r *partnerUpdateRequest) populate(p *model.Partner) {
	r.Partner.Name = p.Name
}

func (r *partnerUpdateRequest) bind(c echo.Context, p *model.Partner) error {
	if err := c.Bind(r); nil != err {
		return err
	}
	if err := c.Validate(r); nil != err {
		return err
	}
	p.Name = r.Partner.Name
	return nil
}
