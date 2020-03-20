package main

import (
	"github.com/kekhuay/echo-installation-service/db"
	"github.com/kekhuay/echo-installation-service/handler"
	"github.com/kekhuay/echo-installation-service/router"
	"github.com/kekhuay/echo-installation-service/store"
)

func main() {
	r := router.New()
	v1 := r.Group("/api/v1")

	db.New()

	ps := store.NewPartnerStore()
	h := handler.NewHandler(ps)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8080"))
}
