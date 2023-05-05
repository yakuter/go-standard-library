package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Starships struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name                 string   `json:"name"`
		Model                string   `json:"model"`
		Manufacturer         string   `json:"manufacturer"`
		CostInCredits        string   `json:"cost_in_credits"`
		Length               string   `json:"length"`
		MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
		Crew                 string   `json:"crew"`
		Passengers           string   `json:"passengers"`
		CargoCapacity        string   `json:"cargo_capacity"`
		Consumables          string   `json:"consumables"`
		HyperdriveRating     string   `json:"hyperdrive_rating"`
		MGLT                 string   `json:"MGLT"`
		StarshipClass        string   `json:"starship_class"`
		Pilots               []string `json:"pilots"`
		Films                []string `json:"films"`
		Created              string   `json:"created"`
		Edited               string   `json:"edited"`
		URL                  string   `json:"url"`
	} `json:"results"`
}

func main() {
	// Read the JSON file into a byte slice
	jsonFile, err := os.ReadFile("starships.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the byte slice into a Starships struct
	var starships Starships
	err = json.Unmarshal(jsonFile, &starships)
	if err != nil {
		panic(err)
	}

	// Create an HTTP handler function to serve the Starships struct as JSON
	http.HandleFunc("/starships", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(starships)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	// Start the server
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
