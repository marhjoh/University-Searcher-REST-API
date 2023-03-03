package main

import (
	"assignment-1/handlers"
	"assignment-1/predefined"
	"assignment-1/uptime"
	"log"
	"net/http"
	"os"
)

// The main function handles ports assignment, sets up handler endpoints and starts the HTTP-server
func main() {
	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = predefined.PORT
	}

	// Set up handler endpoints
	http.HandleFunc(predefined.UNIVERSITYINFORMATION_PATH, handlers.HandlerUniversityInformation)
	http.HandleFunc(predefined.NEIGHBOURUNIVERSITIES_PATH, handlers.HandlerNeighbourUniversities)
	http.HandleFunc(predefined.DIAG_PATH, handlers.HandlerDiag)

	// Starting HTTP-server
	log.Println("Starting server on port " + port + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
