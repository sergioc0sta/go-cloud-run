package main

import (
	"net/http"
	"fmt"

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

	http.ListenAndServe(":8081", mux)
	fmt.Println("Server running on port 8081")
}
