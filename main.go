package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"github.com/WTIFS/tantan-demo/controller"
)

func main() {
	r := mux.NewRouter()
	controller.HandlerUser(r.PathPrefix("/users"))

	serverMsg := http.ListenAndServe(":8000", r)
	log.Fatal(serverMsg)
}