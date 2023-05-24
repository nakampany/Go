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
	var s2 string = "hello"
	var Toint_s int
	var Toint_err int

	Toint_s, _ = strconv.Atoi(s)
	fmt.Println(Toint_s) // int 123

	var err error

	Toint_err, err = strconv.Atoi(s2)
	fmt.Println(Toint_err) // 0
	fmt.Println(err)       // strconv.Atoi: parsing "hello": invalid syntax

	// 文字列 (string) → 論理型 (bool)
	var s3 string = "true"
	var toBool bool

	toBool, _ = strconv.ParseBool(s3)
	fmt.Println(toBool) // ture

}
