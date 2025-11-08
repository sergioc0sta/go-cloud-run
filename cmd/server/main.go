package main

import (
	"net/http"
	"log"

	"github.com/sergioc0sta/go-cloud-run/configs"
	"github.com/sergioc0sta/go-cloud-run/internal/infra/handlers"
	"github.com/sergioc0sta/go-cloud-run/internal/infra/temperature"
)

func main() {
	cfg, _ := configs.LoadConfig(".")

	viaCepAPI := cfg.ViaCepAPI
	tempAPI := cfg.WeatherAPI

	mux := http.NewServeMux()
	temperatureClient := temperature.NewTemperatureClient(viaCepAPI, tempAPI) 

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/temp", handlers.WeatherHandler(temperatureClient))

	log.Println("Server is running on 8080 port...")
	http.ListenAndServe(":8080", mux)
}
