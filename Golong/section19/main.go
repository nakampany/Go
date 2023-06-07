package main

import "fmt"

// 引数の型を[T any](s []T)にすることで、どんな型でも受け取れるようになる
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
}
