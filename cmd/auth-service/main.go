package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"idp_mvp/internal/auth-service/v1/repository"
	"idp_mvp/internal/auth-service/v1/service"
	apiAuth "idp_mvp/pkg/api/generated/auth-service"
	"idp_mvp/pkg/logger"
	"idp_mvp/pkg/middleware"

	_ "idp_mvp/internal/auth-service/v1/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

var SugaredLogger *zap.SugaredLogger

// @title IDP Auth-Service
// @version 1.0
// @description Authentication service for payloads
// @BasePath /
// @contact.name platform-team
// @contact.email platform-team@myorg.io
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	SugaredLogger = logger.Init("auth-service")
	defer SugaredLogger.Sync()
	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	SugaredLogger.Infof("starting the server on %d", 4000)

	go func() {
		if err := CreateHttpServer(); err != nil && err != http.ErrServerClosed {
			SugaredLogger.Fatalf("Failed to start auth-service :%v", err)
			os.Exit(1)
		}
	}()

	<-sigCh
}

func CreateHttpServer() error {
	svc := service.NewUserService(repository.NewInMemoryReps())
	handler := service.New(SugaredLogger, svc)

	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("this is a auth"))
	})

	muxRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4000/swagger/doc.json"),
	)).Methods("GET", "OPTIONS", "POST", "PUT")

	serverOptions := apiAuth.GorillaServerOptions{
		BaseURL: "",
		Middlewares: []apiAuth.MiddlewareFunc{
			middleware.Logging,
			middleware.RecoveryMiddleware,
			middleware.CORSMiddleware,
		},
		BaseRouter: muxRouter,
	}

	r := apiAuth.HandlerWithOptions(handler, serverOptions)

	server := http.Server{
		Handler:      r,
		Addr:         ":4000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
