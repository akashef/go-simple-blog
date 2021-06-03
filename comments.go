package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var myComment Comment
	db.First(&myComment, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&myComment)
}
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	body := processBody(r.Body)
	var myComment Comment
	db.First(&myComment, body["id"])
	db.Model(&myComment).Update("Body", body["Body"])
	json.NewEncoder(w).Encode(&myComment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var myComment Comment
	db.Delete(&myComment, params["id"])
	checkErr(err)
	json.NewEncoder(w).Encode(&myComment)
}
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var myComment Comment
	json.NewDecoder(r.Body).Decode(&myComment)
	db.Create(&myComment)
	json.NewEncoder(w).Encode(&myComment)
}
