package main

import (
	"fmt"
	"net/http"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users endpoint hit")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpoint hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete endpoint")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update")
}
