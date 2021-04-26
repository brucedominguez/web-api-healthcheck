package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brucedominguez/web-api-healthcheck/driver"
	"github.com/brucedominguez/web-api-healthcheck/handler"
	"github.com/gorilla/mux"
)

func main() {
	driver.Init(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	r := mux.NewRouter()

	r.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
