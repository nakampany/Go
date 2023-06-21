package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {
	i := []int{5, 3, 2, 4, 5, 6, 7, 8, 1}
	s := []string{"a", "c", "b"}

	fmt.Println(i, s) // [5 3 2 4 5 6 7 8 1] [a b c]

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s) // [1 2 3 4 5 5 6 7 8] [a b c]

	el := []Entry{
		{"A", 2},
		{"b", 2},
		{"c", 2},
		{"f", 1},
		{"a", 2},
		{"u", 6},
		{"A", 4},
		{"e", 22},
		{"g", 20},
	}

	fmt.Println(el) // [{A 2} {b 2} {c 2} {f 1} {a 2} {u 6} {A 4} {e 22} {g 20}]

	// 2回ソートしたら、アルファベット順がくずれる
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println(el) // [{A 2} {A 4} {a 2} {b 2} {c 2} {e 22} {f 1} {g 20} {u 6}]
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println(el) // [{f 1} {A 2} {a 2} {b 2} {c 2} {A 4} {u 6} {g 20} {e 22}]

	// 2回ソートしても、アルファベット順がくずれない
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	fmt.Println(el)
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println(el) // [{A 2} {A 4} {a 2} {b 2} {c 2} {e 22} {f 1} {g 20} {u 6}]
}
