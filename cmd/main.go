package main

import (
	"fmt"
	"log"
	"net/http"

	"back-eco-go/routes"
)

func main() {
	router := routes.InitRoutes()

	port := "8080"
	fmt.Printf("Server running on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
