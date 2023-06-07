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

// 3
// Underlying Typeのみ許可
type Number2 interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Add2[T Number2](a, b T) T {
	return a + b
}

type MyInt int

// 4
type SampleType interface {
	string | int
}

func sampleFuncGenerics[T SampleType](x T) T {
	return x
}

func main() {
	// 1
	PrintSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	PrintSlice[string]([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
	// 2
	fmt.Println(Add[int](1, 2))
	// 3
	fmt.Println(Add2[MyInt](1, 2))
	// 4
	fmt.Println(sampleFuncGenerics[int](1))
	fmt.Println(sampleFuncGenerics[string]("a"))

}
