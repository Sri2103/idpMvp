package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type handlers struct {
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *handlers {
	return &handlers{
		logger: logger,
	}
}

// healthzHandler godoc
// @Summary health of the server handler
// @Success 200
// @Router /healthz [get]
func (handlers) GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// readyHandler godoc
// @Summary readyness of the server
// @Success 200
// @Router /ready [get]
func (handlers) GetReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
