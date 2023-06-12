package main

import (
	"fmt"
	"log"
	"net/http"
	"udemy/Golong/TodoApp/go/utils"
)

func init() {
	utils.LoggingSettings("./app.log")
}

// メイン関数
func main() {
	log.Println("test")
	fmt.Println("Hello, World!")
	StartMainServer()
}

func StartMainServer() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Path is %s", r.URL.Path[:])
}
