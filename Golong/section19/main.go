package main

import (
	"fmt"
	"strconv"
)

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

// 5
// fは型パラメータを持つ関数
// Tは型パラメータ
// インタフェースStringerは、Tに対する型制約として使われている
func f[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type Stringer interface {
	String() string
}

// MyIntはStringerを実装する
func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

// 6
type Vector[T any] []T

// 7
type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

// 8
type set[T comparable] map[T]struct{}

func (s set[T]) add(x T) {
	s[x] = struct{}{}
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

	// 5
	fmt.Println(f([]MyInt{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	// 6
	v := Vector[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(v)

	// 7
	var t T[int, []*int, *int]
	fmt.Println(t)

	// 8
	s := make(set[int])
	s.add(1)
}
