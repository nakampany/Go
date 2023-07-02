package main

import "fmt"

func main() {
	// スライスの定義

	// スライスの定義: 1
	var s []int
	fmt.Println(s) // []

	// スライスの定義: 2
	s = []int{1, 2, 3}
	fmt.Println(s) // [1 2 3]

	// スライスの定義: 3
	s = make([]int, 3)
	fmt.Println(s) // [0 0 0]

	// スライスの定義: 4
	s = make([]int, 3, 5)
	fmt.Println(s) // [0 0 0]

	// スライスの定義: 5
	s = []int{1: 5, 4: 10}
	fmt.Println(s) // [0 5 0 0 10]

	// 配列の定義： 1
	var a [3]int
	fmt.Println(a) // [0 0 0]

	// 配列の定義： 2
	a = [3]int{1, 2, 3}
	fmt.Println(a) // [1 2 3]

	// 配列の定義： 3
	a = [...]int{1, 2, 3}
	fmt.Println(a) // [1 2 3]

	// mapの定義： 1
	var m map[string]int
	fmt.Println(m) // map[]

	// mapの定義： 2
	m = map[string]int{"x": 10, "y": 20}
	fmt.Println(m) // map[x:10 y:20]

	// mapの定義： 3
	m = make(map[string]int)
	fmt.Println(m) // map[]

	// スライスの要素の取得
	s = []int{1, 2, 3}
	fmt.Println(s[0]) // 1
	fmt.Println(s[1]) // 2
	fmt.Println(s[2]) // 3

	// スライスのコピー
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)

	fmt.Println(n, dst) // 5 [1 2 3 4 5]

	// スライスのコピー: 2
	src = []int{1, 2, 3, 4, 5}
	dst = src[:]
	n = copy(dst, src)

	fmt.Println(n, dst) // 5 [1 2 3 4 5]

	// スライスの結合
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)

	fmt.Println(s3) // [1 2 3 4 5 6]

	// スライスの要素の追加
	s = []int{1, 2, 3, 4, 5}
	i := 2
	s = append(s[:i], append([]int{100}, s[i:]...)...)

	fmt.Println(s) // [1 2 100 3 4 5]

	// スライスの要素の追加: 2
	s = []int{1, 2, 3, 4, 5}
	i = 2
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = 100

	fmt.Println(s) // [1 2 100 3 4 5]

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

	// copy関数を使った要素のシフト
	s = []int{1, 2, 3, 4, 5}
	i = 2
	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0

	fmt.Println(s) // [1 2 4 5 0]

	// スライスの要素をシフト: 2
	s = []int{1, 2, 3, 4, 5}
	i = 2
	s = append(s[:i], s[i:len(s)-1]...)
	s = append(s, 0)

	fmt.Println(s) // [1 2 4 5 0]

	// スライスの要素を偶数でフィルタリング
	s = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); {
		if s[i]%2 == 0 {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(s) // [1 3 5]

	// スライスの要素を偶数でフィルタリング: 2
	s = []int{1, 2, 3, 4, 5}
	n := 0
	for _, v := range s {
		if v%2 != 0 {
			s[n] = v
			n++
		}
	}
	s = s[:n]

	// スライスの要素を奇数でフィルタリング
	s = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); {
		if s[i]%2 != 0 {
			s = append(s[:i], s[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(s) // [2 4]

	// スライスを任意の要素数に分割
	s = []int{1, 2, 3, 4, 5}
	n = 2
	for i := 0; i < len(s); i += n {
		fmt.Println(s[i : i+n])
	}

	// [1 2]
	// [3 4]
	// [5]

	// スライスを任意の要素数に分割: 2
	s = []int{1, 2, 3, 4, 5}
	n = 2
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		fmt.Println(s[i:end])
	}

	// [1 2]
	// [3 4]
	// [5]

	// スライスを任意の要素数に分割: 3
	src1 := []int{1, 2, 3, 4, 5}
	size := 2
	dst1 := make([][]int, 0, (len(src1)+size-1)/size)
	for size < len(src1) {
		src1, dst1 = src1[size:], append(dst1, src1[0:size:size])
	}
	dst1 = append(dst1, src1)

	fmt.Println(dst1) // [[1 2] [3 4] [5]]

}


// スライスをシャッフリングする
src1 = []int{1, 2, 3, 4, 5}
dst1 := make([]int, len(src1))
perm := rand.Perm(len(src1))
for i, v := range perm {
	dst1[v] = src1[i]
}

fmt.Println(dst1) // [3 5 1 4 2]


// スライスをシャッフリングする: 2
s = []int{1, 2, 3, 4, 5}
for i := len(s) - 1; i > 0; i-- {
	j := rand.Intn(i + 1)
	s[i], s[j] = s[j], s[i]
}

fmt.Println(s) // [3 5 1 4 2]


// スライスをバッチ処理する
actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
batchSize := 3
batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

for batchSize < len(actions) {
	actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
}
batches = append(batches, actions)
fmt.Println(batches)

// ２次元スライスの定義：１
var s [][]int

// ２次元スライスの定義：２
s := [][]int{{1, 2}, {3, 4}, {5, 6}}
fmt.Println(s) // [[1 2] [3 4] [5 6]]

// ２次元スライスの定義：３
s := make([][]int, 5)
fmt.Println(s) // [[0] [0] [0] [0] [0]]

// ２次元スライスの定義：4
s := make([][]int, 5, 10)
fmt.Println(s) // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

// ２次元スライスの定義：5
s := make([][]int, 5)
for i := range s {
	s[i] = make([]int, 3)
}

fmt.Println(s) // [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]

// ２次元スライスの要素を取得する
s = [][]int{{1, 2}, {3, 4}, {5, 6}}
fmt.Println(s[1][1]) // 4

// ２次元スライスの要素追加
s = [][]int{{1, 2}, {3, 4}, {5, 6}}
s = append(s, []int{7, 8})
fmt.Println(s) // [[1 2] [3 4] [5 6] [7 8]]

// ２次元スライスの要素削除
s = [][]int{{1, 2}, {3, 4}, {5, 6}}
i := 1
s = append(s[:i], s[i+1:]...)
fmt.Println(s) // [[1 2] [5 6]]

// ２次元スライスの要素削除: 2
s = [][]int{{1, 2}, {3, 4}, {5, 6}}
i := 1
s = s[:i+copy(s[i:], s[i+1:])]
fmt.Println(s) // [[1 2] [5 6]]

// ２次元スライスの要素削除: 3
s = [][]int{{1, 2}, {3, 4}, {5, 6}}
i := 1

// ２次元スライスの要素結合
s1 = [][]int{{1, 2}, {3, 4}, {5, 6}}
s1 = [][]int{{1, 2}, {3, 4}, {5, 6}}
s1 = append(s1, s2...)

fmt.Println(s1) // [[1 2] [3 4] [5 6] [1 2] [3 4] [5 6]]

// ２次元スライスの要素のソート
s = [][]int{{3, 4}, {1, 2}, {5, 6}}
sort.Slice(s, func(i, j int) bool {
	return s[i][0] < s[j][0]
})
fmt.Println(s) // [[1 2] [3 4] [5 6]]

// ２次元スライスの要素のソート: 2
s = [][]int{{3, 4}, {1, 2}, {5, 6}}
sort.SliceStable(s, func(i, j int) bool {
	return s[i][0] < s[j][0]
})
fmt.Println(s) // [[3 4] [1 2] [5 6]]

// ２次元スライスの要素のソート: 3
s = [][]int{{2, 4}, {1, 3}, {5, 6}}
sort.Slice(s, func(i, j int) bool {
	return s[i][1] < s[j][1]
}
fmt.Println(s) // [[1 3] [2 4] [5 6]]




