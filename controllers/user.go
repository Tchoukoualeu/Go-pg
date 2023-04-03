package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id int
	Country string
	FirstName string
	LastName string
	Point int
}


var users = []User{
	{1, "France", "Alain", "Gerard", 102},
	{2, "Belgium", "Madonna", "Dolorean", 302},
	{3, "Germany", "markus", "Miele", 876},
	{4, "Poland", "Arthur", "Oski", 303},
	{5, "Cameroon", "Paul", "Zogo", 301},
}

func ReturnUsers(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func ReturnUsersByCountry(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	country := vars["country"]
	
	usersByCountry := &[]User{}
	
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
	
	usersByCountry := &[]User{}
	
	for _,  user := range users{
		if(user.Id == id){
			*usersByCountry = append(*usersByCountry, user)
		}
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersByCountry)

}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser User
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

	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	fmt.Print(updatedUser)
	for k, u := range users{
		if u.Id == id {
			users = append(users[:k], users[k+1:]...)
			users = append(users, updatedUser)
		}
	}

	fmt.Print(users)

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