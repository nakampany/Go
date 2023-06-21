package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	var User1 User
	fmt.Println(User1) // { 0} -> 初期値が表示される
	User1.Age = 20
	fmt.Println(User1.Age)

	User2 := User{}
	fmt.Println(User2)

	User3 := User{Name: "naka", Age: 20}
	fmt.Println(User3)

	// アドレス演算子を使う
	// ４、５同じ
	// ユースケース
	User4 := new(User) // &{} -> アドレスが表示
	fmt.Println(User4)

	User5 := &User{} // よく使う　関数の引数
	fmt.Println(User5)

	//比較
	User_1 := User{}
	fmt.Println(User_1)
	User_2 := &User{}
	fmt.Println(User_2)

	UpdateUser1(User_1) // { 0}
	UpdateUser2(User_2) // &{ 0}
	fmt.Println(User_1) // { 0}
	fmt.Println(User_2) // &{B 2}

	// UpdateUser1(User_2)
	// UpdateUser2(User_1)
	// cannot use User_2 (*User) as User value in argument to UpdateUser
}

func UpdateUser1(user User) {
	user.Name = "A"
	user.Age = 1
}
func UpdateUser2(user *User) {
	user.Name = "B"
	user.Age = 2
}
