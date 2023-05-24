package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 123
	var toString_s string
	var toString_Unicode string

	toString_s = strconv.Itoa(i)
	toString_Unicode = string(rune(i))
	fmt.Println(i)                // int 123
	fmt.Println(toString_s)       // string "123"
	fmt.Println(toString_Unicode) // 123がrune(Unicode)として認識され "{" になるので注意!!
}
