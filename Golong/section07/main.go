package main

import "fmt"

/*
関数
*/
func p(i interface{}) {
	fmt.Println(i)
}

func add(x int, y int) int {
	return x + y
}

func add2() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	p("It's sample function")
	p("Do you like Golang?")
	fmt.Println(add(42, 13)) // => 55

	// 引数なし
	sayWorld := func() {
		fmt.Println("world")
	}
	sayWorld()

	// 引数
	say1 := func(word string) {
		fmt.Println(word)
	}
	say1("hoge")

	// 即実行
	func(word string) {
		fmt.Println(word)
	}("hogehoge")

	// クロージャー
	// add2は関数が戻り値なので変数に代入
	add2 := add2()
	fmt.Println(add2())
	fmt.Println(add2())
	fmt.Println(add2())
}
