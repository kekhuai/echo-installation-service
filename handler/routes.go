package handler

import "github.com/labstack/echo"

func (h *Handler) Register(v1 *echo.Group) {
	partners := v1.Group("/partners")
	partners.POST("", h.CreatePartner)
}
