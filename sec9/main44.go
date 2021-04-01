package main

import "fmt"

// スライスの要素数は可変長

func main() {
	sl := []int{10, 20}
	fmt.Println(sl)

	sl = append(sl, 30)
	fmt.Println(sl)

	sl = append(sl, 40, 50, 60)
	fmt.Println(sl)

	// スライスの作成
	sl2 := make([]int, 5)
	fmt.Println(sl2)

	// 要素数を数える
	fmt.Println(len(sl2))

	// キャパシティ(容量)を調べる
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(sl3)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

}
