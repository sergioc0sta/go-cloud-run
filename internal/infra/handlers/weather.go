package handlers

import (
	"net/http"
	"strconv"

	"github.com/sergioc0sta/go-cloud-run/internal/validate"
)

type WeatherResponse struct {
	Location string `json:"location"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	idValid := validate.CepValidator(cep)

	if !idValid {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusUnprocessableEntity) 
		w.Write([]byte("invalid zipcode"))
		return
	}

	w.Write([]byte("Weather data for CEP: " + strconv.FormatBool(idValid)))
}
