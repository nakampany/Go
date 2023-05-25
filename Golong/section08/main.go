package main

import (
	"errors"
	"fmt"
)

func main() {
	err := bar()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Println("Normal End.")

	sl := []string{"python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
		fmt.Println(i)
		fmt.Println(v)
	}
	m := map[string]int{"udon": 500, "ra-men": 800, "soba": 700}
	fmt.Println(m) // => map[ra-men:800 soba:700 udon:500]
}

func bar() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()

	foo()
	return
}

func foo() {
	panic(errors.New("Panic!"))
}
