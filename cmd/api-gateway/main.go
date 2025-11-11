package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"idp_mvp/pkg/api/generated/api-gateway"
	"idp_mvp/pkg/logger"
	"idp_mvp/pkg/middleware"

	_ "idp_mvp/internal/api-gateway/v1/docs"
	"idp_mvp/internal/api-gateway/v1/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

var SugaredLogger *zap.SugaredLogger

// @title IDP API Gateway
// @version 1.0
// @description Entry point for all the clients and cli
// @BasePath /api/v1
// @contact.name plateform-team
// @contact.email platform-team@myorg.io
// @license name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// cfg := config.Load("api-gateway")

	// Initialize logger
	SugaredLogger = logger.Init("api-gateway")
	defer SugaredLogger.Sync()

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	SugaredLogger.Infof("starting the server on %d", 5000)

	go func() {
		if err := CreateHttpServer(); err != nil && err != http.ErrServerClosed {
			SugaredLogger.Fatalf("Failed to start API Gateway:%v", err)
			os.Exit(1)
		}
	}()

	<-sigCh
}

func CreateHttpServer() error {
	handler := handlers.New(SugaredLogger)

	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("this is a gateway"))
	})

	muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5000/swagger/doc.json"),
	)).Methods("GET", "OPTIONS", "POST", "PUT")

	serverOptions := api.GorillaServerOptions{
		BaseURL: "/api/v1",
		Middlewares: []api.MiddlewareFunc{
			middleware.Logging,
			middleware.RecoveryMiddleware,
			middleware.CORSMiddleware,
		},
		BaseRouter: muxRouter,
	}

	r := api.HandlerWithOptions(handler, serverOptions)

	server := http.Server{
		Handler:      r,
		Addr:         ":5000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
