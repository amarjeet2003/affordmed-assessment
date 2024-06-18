package main

import (
	"log"
	"net/http"
	"top-products-microservice/routes"
)

func main() {
	router := routes.SetupRouter()

	// Start the server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
