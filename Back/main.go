package main

import (
	"Backend/controllers"
	"Backend/database"
	"log"
	"net/http"
)

func main() {
	database.InitDB("mailService.db")
	// Init
	controllers.Routes()
	log.Println("Port: :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
