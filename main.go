package main

import (
	"goP10/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.ReturnUsers).Methods("GET")
	router.HandleFunc("/users/country/{country}", controllers.ReturnUsersByCountry).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.ReturnUsersById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUserById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}