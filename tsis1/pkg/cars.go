package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Car struct {
	ID        int    `json:"id"`
	Brand     string `json:"brand"`
	Model     string `json:"model"`
	Year      int   `json:"year"`
}

type Response struct {
	Cars []Car `json:"persons"`
}

func prepareResponse() []Car {
	var cars []Car

	cars = append(cars, Car{ID: 1, Brand: "BMW", Model: "M2", Year: 2018})
	cars = append(cars, Car{ID: 2, Brand: "BMW", Model: "M3", Year: 2018})
	cars = append(cars, Car{ID: 3, Brand: "BMW", Model: "M4", Year: 2019})
	cars = append(cars, Car{ID: 4, Brand: "BMW", Model: "M5", Year: 2008})
	cars = append(cars, Car{ID: 5, Brand: "BMW", Model: "M6", Year: 2009})
	cars = append(cars, Car{ID: 6, Brand: "BMW", Model: "M8", Year: 2019})
	
  
	return cars
}

func Cars(w http.ResponseWriter, r *http.Request) {
	var response Response
	cars := prepareResponse()

	response.Cars = cars

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
	 return
	}
  
	w.Write(jsonResponse)
}

func CarByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carIDStr, ok := vars["id"]

	if !ok {
		http.Error(w, "Missing car ID", http.StatusBadRequest)
		return
	}

	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	allCars := prepareResponse()

	var foundCar Car
	for i := 0; i < len(allCars); i++ {
		if allCars[i].ID == carID {
			foundCar = allCars[i]
			break
		}
	}

	if foundCar.ID == 0 {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(foundCar)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}