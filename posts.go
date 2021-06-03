package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var myPost Post
	db.First(&myPost, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&myPost)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	body := processBody(r.Body)
	var myPost Post
	db.First(&myPost, body["id"])
	db.Model(&myPost).Update("Body", body["Body"])
	db.Model(&myPost).Update("Title", body["Title"])
	json.NewEncoder(w).Encode(&myPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var myPost Post
	db.Delete(&myPost, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&myPost)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var myPost Post
	json.NewDecoder(r.Body).Decode(&myPost)
	db.Create(&myPost)
	json.NewEncoder(w).Encode(&myPost)
}
