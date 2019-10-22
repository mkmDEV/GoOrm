package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	router.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	handleRequests()
}
