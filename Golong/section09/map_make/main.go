package main

import (
	"fmt"
)

func main() {
	// make
	x := make([]int, 0)
	fmt.Println(x) // []

	// map
	// makeを使って初期化　　キャパシティは設定できる
	m1 := make(map[string]int, 3)
	fmt.Println(m1) // 	map[]

	// makeを使わずに初期化　　キャパシティは自動で確保
	m2 := map[string]int{}
	fmt.Println(m2) // 	map[]

	var m map[string]int
	// 初期化していない状態で値をセットしようとするとパニックする
	if m == nil {
		m = make(map[string]int)
	}
	m["key"] = 100
	fmt.Println(m)
	fmt.Println(m["key"])
	// map[key:100]
	// 100

}
