package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 入力全体を読み込む
	f, _ := os.Open("./main.go")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs))

	// ファイルに書き込む
	if err := ioutil.WriteFile("./main.go", bs, 06666); err != nil {
		log.Fatalln(err)
	}
}