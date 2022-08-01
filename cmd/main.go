package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rollic/pkg/middleware"
	"rollic/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/login", routes.Login).Methods("POST")
	router.HandleFunc("/user", routes.Add).Methods("PUT")
	router.HandleFunc("/user", middleware.IsAuthorized(routes.Update)).Methods("PATCH")
	router.HandleFunc("/user/{id}", routes.Delete).Methods("DELETE")
	router.HandleFunc("/user/{id}", routes.Get).Methods("GET")
	router.HandleFunc("/users/all", routes.All).Methods("GET")
	router.HandleFunc("/welcome", middleware.IsAuthorized(routes.Welcome)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
