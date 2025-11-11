package main

import (
	"net/http"
	"os"

	"idp_mvp/pkg/logger"
	"idp_mvp/pkg/middleware"

	_ "idp_mvp/internal/api-gateway/v1/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

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
	log := logger.Init("api-gateway")
	defer log.Sync()

	// create mux router
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.Logging)
	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.CORSMiddleware)

	r.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("this is a gateway"))
	})

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5000/swagger/doc.json"),
	)).Methods("GET", "OPTIONS", "POST", "PUT")

	log.Infof("starting the server on %d", 5000)
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("Failed to start API Gateway:%v", err)
		os.Exit(1)
	}
}
