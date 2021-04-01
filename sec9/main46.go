package main

import "fmt"

// スライスのコピー

func main() {
	sl := []int{1, 2, 3, 4, 5}
	var sl2 []int //同型のsl2を作成
	// sl2 := make([]int, 5)

	sl2 = sl     //slをsl2に代入
	sl2[0] = 100 //sl2[0]の値を変更

	fmt.Println(sl2)
	fmt.Println(sl) //sl2での変更によりslの要素も変更された

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5)
	fmt.Println(sl3)
	fmt.Println(sl4)

	n := copy(sl4, sl3)
	sl4[0] = 100

	fmt.Println(sl3) //sl3からsl4をコピーしただけなので、sl3の値は変更されず
	fmt.Println(sl4) // コピー後の値
	fmt.Println(n)   // コピーの要素数

	// makeでスライスを作成
	/*
		make([]T, len, cap)
		make() 関数の第 1 引数 []T が型
		第 2 引数 (len) が 長さ
		第 3 引数 (cap) が 容量
	*/

}
