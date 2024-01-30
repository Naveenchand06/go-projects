package main

import (
	"log"
	"net/http"

	"github.com/Naveenchand06/go-projects/go-mongo-api/config"
	"github.com/Naveenchand06/go-projects/go-mongo-api/routes"
	"github.com/gorilla/mux"
)


func main() {
	config.ConnectToDB()
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":5020", router))
}