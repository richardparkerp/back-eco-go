package main

import (
	"fmt"
	"log"
	"net/http"

	"back-eco-go/config"
	"back-eco-go/routes"
)

func main() {
	log.Println("===Srating server===")
	config.InitDB()
	log.Println("===DB connected===")
	router := routes.InitRoutes()
	log.Println("===Routes succcessfully intialiazed===")
	port := "8080"

	fmt.Printf("Server running on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
