package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// APIResponse struct to store data fetched from the open weather API
// We just need to extract two fields from the response body of GET call - weather.main and main.temp
// We need Weather as a slice of structs because it is possible to get multiple weather conditions for a request.

type APIResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

// TODO: Add your API key here for testing
const ApiKey = "DUMMY_KEY"

const fixedUrl = "https://api.openweathermap.org/data/2.5/weather"

func requestHandler(w http.ResponseWriter, r *http.Request) {

	latitude, longitude := r.URL.Query().Get("lat"), r.URL.Query().Get("lon")
	if latitude == "" || longitude == "" {
		http.Error(w, "Latitude or Longitude is missing", http.StatusBadRequest)
		return
	}

	// Construct the URL
	url := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric", fixedUrl, latitude, longitude, ApiKey)
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Could not call the API", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	// Check for successful response
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		http.Error(w, fmt.Sprintf("GET request failed with error: %s", body), response.StatusCode)
		return
	}

	var weather1 APIResponse

	// Decode the response
	err = json.NewDecoder(response.Body).Decode(&weather1)
	if err != nil {
		http.Error(w, "Error parsing weather data", http.StatusInternalServerError)
		return
	}

	temperature := weather1.Main.Temp
	condition := weather1.Weather[0].Main
	var tempCat string

	// Decide temperature category
	if temperature < 15 {
		tempCat = "Cold"
	} else if temperature < 25 {
		tempCat = "Moderate"
	} else {
		tempCat = "Hot"
	}

	fmt.Fprintf(w, "Temperature is: %s, Weather is: %s", tempCat, condition)
}

func main() {
	http.HandleFunc("/weather", requestHandler)

	var port string

	checkPort := os.Getenv("PORT")
	if checkPort != "" {
		port = checkPort
	} else {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
