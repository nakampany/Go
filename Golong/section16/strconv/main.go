package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b string = "123"

	// ２valueが必要
	a, _ = strconv.Atoi(b /* string */)
	fmt.Println(a)

	var i int = 123
	var s string

	s = strconv.Itoa(i /* int */)
	fmt.Println(s)
}
