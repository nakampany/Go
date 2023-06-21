package main

import (
	"fmt"
)

/*
定数
*/

const (
	aa int = iota
	bb
	cc
)

func main() {
	fmt.Println(aa, bb, cc) // 0, 1, 2

	const a = 1
	var i1 int = a
	var i2 int64 = a

	fmt.Println(i1) //i1 === i2
	fmt.Println(i2)

	const b int = 1
	var i11 int = b
	// エラー
	// ./prog.go:15:18: cannot use b (constant 1 of type int) as int64 value in variable declaration
	// var i22 int64 = b

	fmt.Println(i11) // i1 !== i2
	// fmt.Println(i22)
}
