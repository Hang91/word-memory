package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var dao UsersDAO

// go run *.go
func main() {
	router := mux.NewRouter()
	CreateRouter(router)
	dao = UsersDAO{
		Server:   "localhost:27017",
		Database: "word_memory",
	}

	dao.Connect()

	// start server
	log.Println("Server is listenning port :4000")

	// solves cross origin access issue
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	server := &http.Server{
		Handler:      handler,
		Addr:         "localhost:4000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
