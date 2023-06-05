package controllers

import (
	"log"
	"net/http"
	"udemy/Golong/section18/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
}

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "signup")

	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/", 302)

}
