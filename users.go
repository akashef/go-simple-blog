package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var myUser User
	db.First(&myUser, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&myUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	body := processBody(r.Body)
	var user User
	db.First(&user, body["id"])
	db.Model(&user).Update("name", body["name"])
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.Delete(&user, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&user)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(&user)
}
