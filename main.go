package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var dao UsersDAO

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
	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:4000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
