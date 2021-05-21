package handler

import (
	"app/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/items", h.createItem).Methods(http.MethodPost)
	r.HandleFunc("/api/items", h.getAllItems).Methods(http.MethodGet)
	r.HandleFunc("/api/items/{id:[0-9]+}", h.getItemById).Methods(http.MethodGet)
	r.HandleFunc("/api/items/{id:[0-9]+}", h.updateItem).Methods(http.MethodPut)
	r.HandleFunc("/api/items/{id:[0-9]+}", h.deleteItem).Methods(http.MethodDelete)
	r.HandleFunc("/api/items/{id:[0-9]+}/do", h.doItem).Methods(http.MethodGet)
	r.HandleFunc("/api/items/{id:[0-9]+}/undo", h.undoItem).Methods(http.MethodGet)

	r.Use(h.setJSONMiddleware)

	return r
}
