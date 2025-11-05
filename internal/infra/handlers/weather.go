package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sergioc0sta/go-cloud-run/configs"
	"github.com/sergioc0sta/go-cloud-run/internal/validate"
)

type WeatherResponse struct {
	Location string `json:"localidade"`
}

func fetchWeatherByCEP(cep string) (*WeatherResponse, error) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external service status: %d", resp.StatusCode)
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}
	fmt.Printf("Configs: %+v\n", configs)

	cep := r.URL.Query().Get("cep")
	idValid := validate.CepValidator(cep)

	if !idValid {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zipcode"))
		return
	}

	data, err := fetchWeatherByCEP(cep)
	if err != nil {
		http.Error(w, "failed to call external service", http.StatusBadGateway)
		return
	}

	fmt.Printf("CEP data: %+v\n", data)

	w.Write([]byte("Weather data for CEP: " + strconv.FormatBool(idValid)))
}
