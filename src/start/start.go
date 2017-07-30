package main

import (
	"log"
	"net/http"
	"start/routes"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}
