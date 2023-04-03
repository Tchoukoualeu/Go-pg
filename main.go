package main

import (
	controller "goP10/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controller.ReturnUsers).Methods("GET")
	router.HandleFunc("/users/country/{country}", controller.ReturnUsersByCountry).Methods("GET")
	router.HandleFunc("/users/{id}", controller.ReturnUsersById).Methods("GET")
	router.HandleFunc("/users/{id}", controller.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.DeleteUserById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}