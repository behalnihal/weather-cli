package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)


type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		}`json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				}`json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		}`json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	api_key := os.Getenv("API_KEY")

	q := "Gurgaon"

	if len(os.Args) >= 2{
		q = os.Args[1];
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + api_key + "&q=" + q + "&aqi=yes");
	if err != nil{
		panic(err);
	}
	defer res.Body.Close();
	
	if res.StatusCode != 200 {
		panic("API not available")
	}

	body, err := io.ReadAll(res.Body);
	if err != nil{
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hour := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %0.fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hour {
		date := time.Unix(hour.TimeEpoch, 0) 
		if date.Before(time.Now()){
			continue;
		}
		fmt.Printf(
			"%s - %0.fC, %0.f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		
	}
}