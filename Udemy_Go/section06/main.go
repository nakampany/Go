package main

import "fmt"

/*
比較演算子
*/
func main() {
	type Point struct {
		X int
		Y int
	}
	m1 := map[Point]int{{X: 100, Y: 200}: 20}
	fmt.Println(m1)

	m2 := map[[2]int]int{{100, 200}: 30}
	fmt.Println(m2)

	var s []int
	var m map[any]any
	var f func()
	fmt.Println(s == nil, m == nil, f == nil) // OK
	// invalid operation: s == s (slice can only be compared to nil)
	// invalid operation: m == m (map can only be compared to nil)
	// invalid operation: f == f (func can only be compared to nil)
	// fmt.Println(s == s, m == m, f == f)       // NG
}
