package main

import (
	"flag"
	"fmt"
)

func main() {
	// コマンドラインのオプション コマンドライン引数を扱う

	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	// 	go run sample1_1.go a b c
	// [a b c]

	flag.Parse()
	fmt.Println(flag.Arg(0), flag.Arg(1))
	// go run sample1_2.go hoge fuga
	// hoge fuga

	// 型名(),型名Var()でフラグの定義をした後,
	// Parse()でそれぞれの変数に取得できます。
	var (
		i = flag.Int("int", 0, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", false, "bool flag")
	)
	flag.Parse()
	fmt.Println(*i, *s, *b)
}
