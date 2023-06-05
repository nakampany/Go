package controllers

import (
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "signup")
}

func login(w *http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}
