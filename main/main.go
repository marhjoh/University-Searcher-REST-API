package main

import (
	"assignment-1/cmd/handlers"
	"assignment-1/cmd/handlers/constants"
	"assignment-1/uptime"
	"log"
	"net/http"
	"os"
)

func main() {
	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = constants.PORT
	}

	// Set up handler endpoints
	http.HandleFunc(constants.UNIVERSITYINFORMATION_PATH, handlers.HandlerUniversityInformation)
	http.HandleFunc(constants.NEIGHBOURUNIVERSITIES_PATH, handlers.HandlerNeighbourUniversities)
	http.HandleFunc(constants.DIAG_PATH, handlers.HandlerDiag)

	// Starting HTTP-server
	log.Println("Starting server on port " + port + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
