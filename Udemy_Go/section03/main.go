package main

import "fmt"

// aaa := 100
var bbb int = 200

func main() {
	var i int = 100
	var s string = "Hello World"
	fmt.Println(i, s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	xi := 300
	xs := "Golang"
	xt, xf := true, false
	fmt.Println(xi, xs, xt, xf)

	// ./main.go:34:5: no new variables on left side of :=
	// ./main.go:35:5: no new variables on left side of :=
	// xi := 400
	// xs := "Python"
	// fmt.Println(xi, xs, xt, xf)

	// ./main.go:40:7: cannot use "Golang" (untyped string constant) as int value in assignment
	// xi = "Golang"

	// ./main.go:5:1: syntax error: non-declaration statement outside function body
	// fmt.Println(aaa)
	// fmt.Println(bbb)

	foo()
}

func foo() {
	fmt.Println("foo")
}
