package main

import (
	"fmt"
	"goPG/controllers"
	"log"
	"net/http"

	// "os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)



func main(){

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.ReturnUsers).Methods("GET")
	router.HandleFunc("/users/country/{country}", controllers.ReturnUsersByCountry).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.ReturnUsersById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUserById).Methods("DELETE")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env could not be loaded")
	}

	// apiKey := os.Getenv("API_KEY")

	// fmt.Print(apiKey)

	fmt.Print("App listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}



