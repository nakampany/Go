package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SayName1(name string) {
	u.Name = name
}
func (u *User) SayName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "naka"}
	fmt.Println(user1.Name)

	// あたい渡
	user1.SayName1("sssss")
	fmt.Println(user1.Name) //naka

	// 参照渡
	user1.SayName2("sssss")
	fmt.Println(user1.Name) // sssss
}
