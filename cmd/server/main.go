package main

import (
	"fmt"
	"net/http"

	"github.com/sergioc0sta/go-cloud-run/configs"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("The WeatherAPI %s \n", configs.WeatherAPI)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8081", mux)
}
