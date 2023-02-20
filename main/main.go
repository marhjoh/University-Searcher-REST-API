package main

import (
	constants "assignment1"
	"assignment1/constants"
	"assignment1/handlers"
	"assignment1/uptime"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = constants.PORT
	}

	http.HandleFunc(constants.UNIINFO_PATH, handlers.HandlerUniInfo)
	http.HandleFunc(constants.NEIGHBOURUNIS_PATH, handlers.HandlerNeighbourUnis)
	http.HandleFunc(constants.DIAG_PATH, handlers.HandlerDiag)

	// Starting HTTP-server
	log.Println("Starting server on port " + port + " ...")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
