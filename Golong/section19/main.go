package main

import "fmt"

// 1
// 引数の型を[T any](s []T)にすることで、どんな型でも受け取れるようになる
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// 2
type Number interface {
	int | int32 | int64 | float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

type Number2 interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Add2[T Number2](a, b T) T {
	return a + b
}

type MyInt int

func main() {
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})

	fmt.Println(Add[int](1, 2))

	fmt.Println(Add2[MyInt](1, 2))

}
