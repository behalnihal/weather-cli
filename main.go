package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load(".env");
	if e != nil{
		log.Fatalf("error loading .env file: %s", e)
	}
	fmt.Println("hello")
	var city string
	apiKey := os.Getenv("API_KEY")
	fmt.Print("Enter city name: ")
	fmt.Scan(&city)
	api := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey;
	response, err := http.Get(api)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err !=  nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}