package handler

import "github.com/kekhuay/echo-installation-service/partner"

// Handler type
type Handler struct {
	partnerStore partner.Store
}

// NewHandler create new instance of handler
func NewHandler(ps partner.Store) *Handler {
	return &Handler{
		partnerStore: ps,
	}
}
