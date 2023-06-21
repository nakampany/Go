package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)

	// 大量の比較については上の表現は向いていない
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)

	re2 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := "0120-111-111"
	ms := re2.FindAllStringSubmatch(s, -1)
	fmt.Println(ms)

	match = re2.MatchString("0120-111-111")
	fmt.Println(match)
}
