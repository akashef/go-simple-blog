package main

import "github.com/gorilla/mux"

func routerHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/create", CreateUser).Methods("POST")
	router.HandleFunc("/user/update", UpdateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/post/{id}", GetPost).Methods("GET")
	router.HandleFunc("/post/create", CreatePost).Methods("POST")
	router.HandleFunc("/post/update", UpdatePost).Methods("POST")
	router.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")

	router.HandleFunc("/comment/{id}", GetComment).Methods("GET")
	router.HandleFunc("/comment/add", CreateComment).Methods("POST")
	router.HandleFunc("/comment/upadte", UpdateComment).Methods("POST")
	router.HandleFunc("/comment/{id}", DeleteComment).Methods("DELETE")
	return router
}
