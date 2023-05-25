// 　ポインタ利用
package main

import "fmt"

func Double2(i *int) {
	*i = *i * 2
}

func main() {

	var n int = 100

	var p *int = &n

	fmt.Println(*p)
	fmt.Println(p)
	Double2(&n)
	fmt.Println(*p)
	fmt.Println(p)
	// 100
	// 0xc0000b2000
	// 200
	// 0xc0000b2000
}
