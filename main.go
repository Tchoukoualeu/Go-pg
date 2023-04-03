package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", Controller.ReturnUsers).Methods("GET")
	router.HandleFunc("/users/country/{country}", Controller.ReturnUsersByCountry).Methods("GET")
	router.HandleFunc("/users/{id}", Controller.ReturnUsersById).Methods("GET")
	router.HandleFunc("/users/{id}", Controller.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", Controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", Controller.DeleteUserById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}