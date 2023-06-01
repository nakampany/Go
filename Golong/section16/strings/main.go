package main

import (
	"fmt"
	"strings"
)

func main() {
	//文字列の結合
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1) // A,B,C
	fmt.Println(s2) // ABC

	// 文字列に含まれる部分文字を検索
	// int
	i1 := strings.Index("ABCDE", "A")         // index番号が返される
	i2 := strings.LastIndex("ABCDABCD", "BC") // 最後に一致したもの、index番号が返される
	i3 := strings.IndexAny("ABCDEA", "A")     // 最初に一致したもの、index番号が返される
	fmt.Println(i1, i2, i3)

	// bool
	b1 := strings.HasPrefix("ABC", "A") // true Aから始まるかどうか
	b2 := strings.HasSuffix("ABC", "C") // true Cでおわるかどうか
	fmt.Println(b1, b2)

	b3 := strings.Contains("ABC", "A")    // true Aが含まれているか
	b4 := strings.ContainsAny("ABC", "C") // true Cが含まれているか
	fmt.Println(b3, b4)

	i6 := strings.Count("ABC", "A") // １ Aが何回出現するか
	i7 := strings.Count("ABC", "")  // true ””（空文字）は文字列＋１が返ってくる
	fmt.Println(i6, i7)

	// 文字列の置換
	i8 := strings.Replace("ABCAAAA", "A", "B", 1)  // はじめから１つ目を AをBに置換
	i9 := strings.Replace("ABCAAAA", "A", "B", -1) // 全ての文字について該当箇所全て AをBに置換
	fmt.Println(i8, i9)

	// 文字列の置換
	i10 := strings.ToLower("ABCAAAA")    // 小文字へ
	i11 := strings.ToUpper("aabbbaaaac") // 大文字へ
	fmt.Println(i10, i11)

	// スペースを消す
	i12 := strings.TrimSpace(" A B C   A   A A A ")
	i13 := strings.TrimSpace("    aabbbaaaac     ")
	fmt.Println(i12, i13)

	//
	i141 := strings.Fields("A B C A A A A")
	i142 := strings.Fields("ABCAAAA")
	i15 := "aabbbaaaac"
	fmt.Println(i141, i142, i15)
	fmt.Println(i141[0]) // A
	fmt.Println(i142[0]) // ABCAAAA
	fmt.Println(i15[0])  // 97

}
