package main

import (
	"fmt"
)

func main() {
	gen1, gen2 := make(chan int), make(chan int)

	// // goルーチンを立てて、gen1やgen2に送信したりする

	// if n1, ok := <-gen1; ok {
	// 	// 処理1
	// 	fmt.Println(n1)
	// } else if n2, ok := <-gen2; ok {
	// 	// 処理2
	// 	fmt.Println(n2)
	// } else {
	// 	// 例外処理
	// 	fmt.Println("neither cannot use")
	// }

	select {
	case num := <-gen1: // gen1から受信できるとき
		fmt.Println(num)
	case num := <-gen2: // gen2から受信できるとき
		fmt.Println(num)
	default: // どっちも受信できないとき
		fmt.Println("neither chan cannot use")
	}
}
