package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"idp_mvp/pkg/logger"

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
// @host  localhost:4000
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// cfg := config.Load("api-gateway")

	// Initialize logger
	SugaredLogger = logger.Init("api-gateway")
	defer SugaredLogger.Sync()

	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	SugaredLogger.Infof("starting the server on %d", 4000)

	go func() {
		if err := CreateHttpServer(); err != nil && err != http.ErrServerClosed {
			SugaredLogger.Fatalf("Failed to start API Gateway:%v", err)
			os.Exit(1)
		}
	}()

	<-sigCh
}

func CreateHttpServer() error {
	hdls := handlers.New(SugaredLogger)

	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4000/swagger/doc.json"),
	)).Methods("GET", "OPTIONS", "POST", "PUT")

	hdls.RegisterRoutes(muxRouter)

	server := http.Server{
		Handler:      muxRouter,
		Addr:         ":4000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
