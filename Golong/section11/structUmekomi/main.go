package main

import "fmt"

type T struct {
	user1 User
}

type User struct {
	Name string
	Age  int
}

func main() {
	t := T{user1: User{Name: "a", Age: 1}}
	fmt.Println(t.user1.Name)

}
