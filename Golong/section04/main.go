package main

import (
	"fmt"
	"strconv"
)

/*
型
*/
func main() {
	//
	var i int = 123
	var toString_s string
	var toString_Unicode string

	toString_s = strconv.Itoa(i)
	toString_Unicode = string(rune(i))
	fmt.Println(i)                // int 123
	fmt.Println(toString_s)       // string "123"
	fmt.Println(toString_Unicode) // 123がrune(Unicode)として認識され "{" になるので注意!!

	// 文字列 (String) → 数値 (int)
	var s string = "123"
	var s2 string = "Draemon"
	var Toint_s int
	var Toint_err int

	Toint_s, _ = strconv.Atoi(s)
	fmt.Println(Toint_s) // int 123

	var err error

	Toint_err, _ = strconv.Atoi(s2)
	fmt.Println(Toint_err)
	fmt.Println(err)
}
