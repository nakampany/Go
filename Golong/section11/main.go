package main

import "fmt"

// Person構造体の定義
type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// Person型の変数の宣言と初期化
	person1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	// Person型のフィールドへのアクセスと出力
	fmt.Println("First name:", person1.firstName)
	fmt.Println("Last name:", person1.lastName)
	fmt.Println("Age:", person1.age)
}

