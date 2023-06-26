package main

import "fmt"

func main() {
	// スライスのコピー
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)

	fmt.Println(n, dst) // 5 [1 2 3 4 5]

	// スライスの結合
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)

	fmt.Println(s3) // [1 2 3 4 5 6]

	// スライスの要素削除
	s := []int{1, 2, 3, 4, 5}
	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s) // [1 2 4 5]

	// copy関数を使った要素削除
	s = []int{1, 2, 3, 4, 5}
	i = 2
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]

	fmt.Println(s) // [1 2 4 5]

	// 順序を保持しない要素削除
	s = []int{1, 2, 3, 4, 5}
	i = 2
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]

	fmt.Println(s) // [1 2 5 4]

	// スライスを逆順にする: 1
	s = []int{1, 2, 3, 4, 5}
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}

	fmt.Println(s) // [5 4 3 2 1]

	// スライスを逆順にする: 2
	s = []int{1, 2, 3, 4, 5}
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}

	fmt.Println(s) // [5 4 3 2 1]

	// スライスを逆順にする: 3
	s = []int{1, 2, 3, 4, 5}
	for i := range s {
		opp := len(s) - 1 - i
		if i >= opp {
			break
		}
		s[i], s[opp] = s[opp], s[i]
	}

	fmt.Println(s) // [5 4 3 2 1]

	// スライスの要素をシフト
	s = []int{1, 2, 3, 4, 5}
	i = 2
	s = append(s[:i], s[i+1:]...)
	s = append(s, 0)

	fmt.Println(s) // [1 2 4 5 0]

}
