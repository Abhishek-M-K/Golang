package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/Abhishek-M-K/mongodb/routes"
)

func main() {
	fmt.Println("MongoDB GoLang API")
	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080")
}
