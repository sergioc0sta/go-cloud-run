package main

import (
	"net/http"

	"github.com/sergioc0sta/go-cloud-run/internal/infra/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/a", handlers.WeatherHandler)

	http.ListenAndServe(":8081", mux)
}
