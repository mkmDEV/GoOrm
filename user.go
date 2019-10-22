package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete endpoint")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update")
}
