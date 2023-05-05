package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StarshipsResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []StarshipInfo `json:"results"`
}

type StarshipInfo struct {
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
	Url                  string   `json:"url"`
}

func main() {
	resp, err := http.Get("https://swapi.dev/api/starships")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var starships StarshipsResponse
	err = json.NewDecoder(resp.Body).Decode(&starships)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	fmt.Printf("Starships count: %d\n\n", starships.Count)
	for _, starship := range starships.Results {
		fmt.Printf("Name: %s\nModel: %s\nManufacturer: %s\nCost in credits: %s\nLength: %s\nMax atmosphering speed: %s\nCrew: %s\nPassengers: %s\nCargo capacity: %s\nConsumables: %s\nHyperdrive rating: %s\nMGLT: %s\nStarship class: %s\nPilots: %v\nFilms: %v\nCreated: %s\nEdited: %s\nURL: %s\n\n",
			starship.Name, starship.Model, starship.Manufacturer, starship.CostInCredits, starship.Length, starship.MaxAtmospheringSpeed, starship.Crew, starship.Passengers, starship.CargoCapacity, starship.Consumables, starship.HyperdriveRating, starship.MGLT, starship.StarshipClass, starship.Pilots, starship.Films, starship.Created, starship.Edited, starship.Url)
	}
}
