package main

import (
	"fmt"
)

func main() {
	// var ch chan int

	// 受信専用のチャネル
	// var ch2 <-chan int
	// 送信専用のチャネル
	// var ch3 chan<- int

	// ch = make(chan int) // cap:0
	// ch2 := make(chan int)    // 0
	ch3 := make(chan int, 5) // 5

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	fmt.Println("len", len(ch3)) // データ数：4

	i := <-ch3
	fmt.Println(i)               //1
	fmt.Println("len", len(ch3)) // データ数：3
	i2 := <-ch3
	fmt.Println(i2)              //2
	fmt.Println("len", len(ch3)) // データ数：2

	fmt.Println(<-ch3)           //3
	fmt.Println("len", len(ch3)) // データ数：2
}
