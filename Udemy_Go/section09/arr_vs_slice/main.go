// 　ポインタ利用
package main

import "fmt"

func Sum(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// スライス
func main() {
	// 配列 要素数指定
	var arr_1 [3]int
	var arr_2 = [3]int{1, 2, 3}
	arr_3 := [3]int{1, 2, 3}
	arr_4 := [...]int{1, 2, 3, 4}
	fmt.Println(arr_1)
	fmt.Println(arr_2)
	fmt.Println(arr_3)
	fmt.Println(arr_4)
	// [0 0 0]
	// [1 2 3]
	// [1 2 3]
	// [1 2 3 4]

	fmt.Println(arr_4[1:]) // [2 3 4]
	// appendできない
	// arr_4 = append(arr_4, 5, 6, 7)
	for _, v := range arr_4 {
		fmt.Println(v)
	}

	// スライス
	var sl_1 []int
	var sl_2 []int = []int{100, 200, 300, 400}
	sl_3 := make([]int, 2)
	fmt.Println(sl_1)
	fmt.Println(sl_2)
	fmt.Println(sl_3)
	// []
	// [100 200]
	// [0 0]

	fmt.Println(sl_2[1:]) // [200 300 400]
	// WARN： 要領以上の要素が追加されるとメモリの消費が倍になる。
	sl_2 = append(sl_2, 500, 600, 700)

	for _, v := range sl_2 {
		fmt.Println(v)
	}
	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))

}
