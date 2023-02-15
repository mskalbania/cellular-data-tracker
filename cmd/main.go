package main

import (
	"cellular-data-tracker/server"
	"flag"
	"os"
	"strings"
)

func main() {
	var config server.Configuration
	flag.IntVar(&config.Port, "port", 8080, "Server port")

	config.Database.ConnectionUrl = os.Getenv("PSQL_CONNECTION_STRING")
	config.DataProvider.Email = os.Getenv("PROVIDER_EMAIL")
	config.DataProvider.Password = os.Getenv("PROVIDER_PASSWORD")
	config.DataProvider.ApiToken = os.Getenv("API_TOKEN")
	config.DataProvider.Numbers = strings.Split(os.Getenv("NUMBERS"), ",")

	server.Start(&config)
}
