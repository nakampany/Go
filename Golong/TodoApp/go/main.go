package main

import (
	"fmt"
	"log"
	"net/http"
	"udemy/Golong/TodoApp/go/app/models"
	"udemy/Golong/TodoApp/go/utils"
)

func init() {
	utils.LoggingSettings("./app.log")
}

// メイン関数
func main() {
	log.Println("Hello World")

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@test.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
	log.Println("Create User")

	StartMainServer()
}

func StartMainServer() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Path is %s", r.URL.Path[:])
}
