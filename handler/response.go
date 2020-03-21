package handler

import (
	"github.com/kekhuay/echo-installation-service/model"
	"github.com/labstack/echo"
)

type partnerResponse struct {
	Name string `json:"name"`
}

type singlePartnerResponse struct {
	Partner *partnerResponse `json:"partner"`
}

type partnerListResponse struct {
	Partners     []*partnerResponse `json:"partners"`
	PartnerCount int64              `json:"partners_count"`
}

func newPartnerResponse(c echo.Context, p *model.Partner) *singlePartnerResponse {
	pr := new(partnerResponse)
	pr.Name = p.Name
	return &singlePartnerResponse{pr}
}

func newPartnerListResponse(partners []model.Partner, count int64) *partnerListResponse {
	r := new(partnerListResponse)
	r.Partners = make([]*partnerResponse, 0)
	for _, p := range partners {
		pr := new(partnerResponse)
		pr.Name = p.Name
		r.Partners = append(r.Partners, pr)
	}
	r.PartnerCount = count
	return r
}
