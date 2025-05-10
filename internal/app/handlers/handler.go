package handlers

import (
	"github.com/go-chi/chi/v5"
	"taskEffectiveMobile/internal/app"
	"taskEffectiveMobile/internal/app/utils"
)

type Handler struct {
	DI *app.DI
}

func InitRoutes(router chi.Router, di *app.DI) {
	handler := Handler{DI: di}
	router.Post("/person", utils.MakeHttpHandler(handler.HandleCreatePerson))
	router.Get("/person", utils.MakeHttpHandler(handler.HandleGetPeople))
	router.Get("/person/{id}", utils.MakeHttpHandler(handler.HandleGetPersonByID))
	router.Put("/person/{id}", utils.MakeHttpHandler(handler.HandleUpdatePerson))
	router.Delete("/person/{id}", utils.MakeHttpHandler(handler.HandleDeletePerson))
}
