package controllers

import (
	"encoding/json"
	"fmt"
	"goPG/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var users = []models.User{
	{1, "France", "Alain", "Gerard", 102},
	{2, "Belgium", "Madonna", "Dolorean", 302},
	{3, "Germany", "markus", "Miele", 876},
	{4, "Poland", "Arthur", "Oski", 303},
	{5, "Cameroon", "Paul", "Zogo", 301},
}

// Not used
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
// end of not used

func ReturnUsers(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func ReturnUsersByCountry(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	country := vars["country"]
	
	usersByCountry := &[]models.User{}
	
	for _,  user := range users{
		if(user.Country == country){
			*usersByCountry = append(*usersByCountry, user)
		}
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersByCountry)

}

func ReturnUsersById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if(err != nil){
		fmt.Print("Unable to convert to string")
	} 
	
	usersByCountry := &[]models.User{}
	
	for _,  user := range users{
		if(user.Id == id){
			*usersByCountry = append(*usersByCountry, user)
		}
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersByCountry)

}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser models.User

	newUser.Id = rand.Intn(100)

	json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Print("Unable to convert to string")
	} 

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	for k, u := range users{
		if u.Id == id {
			users = append(users[:k], users[k+1:]...)
			users = append(users, updatedUser)
		}
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func DeleteUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Print("Unable to convert to string")
	} 

	for k, u := range users{
		if u.Id == id {
			users = append(users[:k], users[k+1:]...)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}