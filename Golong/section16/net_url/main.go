package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://jsonplaceholder.typicode.com/users/1")
	fmt.Println(u.Scheme)   // https
	fmt.Println(u.Host)     // jsonplaceholder.typicode.com
	fmt.Println(u.Path)     // /users/1
	fmt.Println(u.RawQuery) //
	fmt.Println(u.Fragment) //
	fmt.Println(u.Query())  // map[]

	url := &url.URL{}
	url.Scheme = "http"
	url.Host = "google"
	q := url.Query()
	q.Set("q", "Golang")

	url.RawQuery = q.Encode()
	fmt.Println(url) // http://google?q=Golang
}
