package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// コンストラクタ　大事！！！！
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("a", 1)
	fmt.Println(user1)  // &{a 1}
	fmt.Println(*user1) //{a 1}
}
