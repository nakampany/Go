package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	// ランダム占い時間
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num
}

func main() {
	fmt.Println("what is today's lucky number?")

	c := make(chan int) // make(chan int)でチャネルを作成 → getLuckyNum関数に引数として渡す
	go getLuckyNum(c)   // getLuckyNum関数内で得たラッキーナンバーを、チャネルcに送信(c <- num)

	num := <-c // メインゴールーチンで、チャネルcからラッキーナンバーを受信(num := <-c)

	fmt.Printf("Today's your lucky number is %d!\n", num)

	// 使い終わったチャネルはcloseする
	close(c)
}
