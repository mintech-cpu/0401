package main

import "fmt"

func main() {

	//明示的
	var sl []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sl) // -> [1 2 3 4 5]
	//暗示的
	sl10 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl10) // -> [1 2 3 4 5]

	//slの最初の要素を取得してください
	fmt.Println(sl[0]) // -> 1
	//slの2番目から3番目の要素を取得してください
	fmt.Println(sl[1:3]) // -> [2 3]
	//最後の要素を取得してください
	fmt.Println(sl[len(sl)-1]) // -> 5
	//slの最初と最後の要素以外すべてを取得してください
	fmt.Println(sl[1 : len(sl)-1]) // -> [2 3 4]

	// slの要素数を数えてください
	fmt.Println(len(sl)) // -> 5
	// slのキャパシティ(容量)を調べてください
	fmt.Println(cap(sl)) // -> 5
	// slに要素[6 7]を追加してください
	fmt.Println(append(sl, 6, 7))

	// 代入
	sl2 := []string{"apple", "banana", "orange"}
	sl3 := make([]string, 3)
	sl3 = sl2

	// sl3の最初の要素を"peach"に変更し、sl2・sl3を実行
	sl3[0] = "peach"
	fmt.Println(sl2)
	fmt.Println(sl3)

	// コピー
	sl4 := []string{"apple", "banana", "orange"}
	sl5 := make([]string, 3)
	n := copy(sl5, sl4)

	// sl5の最初の要素を"peach"に変更し、sl4・sl5を実行
	sl5[0] = "peach"

	fmt.Println(sl4)
	fmt.Println(sl5)
	fmt.Println(n)

}
