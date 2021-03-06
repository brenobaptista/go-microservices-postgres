package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brenobaptista/go-microservices-postgres/internal/api"
	"github.com/brenobaptista/go-microservices-postgres/internal/db"
)

func main() {
	pgdb, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	router := api.NewAPI(pgdb)

	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router: %v\n", err)
	}
}
