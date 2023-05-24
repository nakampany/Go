package main

import "fmt"

/*
関数
*/
func p(i interface{}) {
	fmt.Println(i)
}

// func add(x, y int) int { ともかける
func add(x int, y int) int {
	return x + y
}

func main() {
	p("It's sample function")
	p("Do you like Golang?")
	fmt.Println(add(42, 13)) // => 55
}