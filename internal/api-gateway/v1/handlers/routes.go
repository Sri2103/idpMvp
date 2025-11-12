package handlers

import (
	"idp_mvp/pkg/middleware"

	"github.com/gorilla/mux"
)

func (h *handlers) RegisterRoutes(r *mux.Router) {
	r.Use(middleware.Logging)
	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.CORSMiddleware)

	v1Router := r.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/healthz", h.GetHealthz)
	v1Router.HandleFunc("/ready", h.GetReady)
}
