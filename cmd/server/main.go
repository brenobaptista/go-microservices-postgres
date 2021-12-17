package main

import (
	"fmt"
	"go-microservices-postgres/pkg/api"
	"go-microservices-postgres/pkg/db"
	"log"
	"net/http"
	"os"
)

func main() {
	_, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	router := api.NewAPI()

	log.Print("we are up and running!")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router: %v\n", err)
	}
}
