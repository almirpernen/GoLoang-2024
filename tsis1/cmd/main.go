package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"tsis1/pkg"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", pkg.HealthCheck).Methods("GET")
	router.HandleFunc("/cars", pkg.Cars).Methods("GET")
	router.HandleFunc("/cars/{id:[0-6]+}", pkg.CarByID).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
