package main

import "fmt"

// スライス
/*
固定長である配列と異なり、スライスは可変長であるため柔軟にデータ（要素）を格納することが可能
*/

func main() {

	// 配列
	var arr [2]int = [2]int{100, 200}
	fmt.Println(arr)
	arr2 := [2]int{100, 200}
	fmt.Println(arr2)

	// スライス
	var sl []int // 配列の宣言と違って、スライスは[]の中に大きさを指定しない
	fmt.Println(sl)

	// スライスの書き方は三通り
	// 明示的
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)
	// 暗示的
	sl3 := []int{100, 200}
	fmt.Println(sl3)
	// makeの使用
	sl4 := make([]int, 2)
	fmt.Println(sl4)
	sl4[0] = 100
	sl4[1] = 200
	fmt.Println(sl4)

	// makeでスライスを作成し、実行結果が[Python GO PHP]になるように
	slp3 := make([]string, 3)
	slp3[0] = "Python"
	slp3[1] = "GO"
	slp3[2] = "PHP"
	fmt.Println(slp3)

	// 要素の取得
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)
	// 最初の要素
	fmt.Println(sl5[0])
	// 最後の要素
	fmt.Println(sl5[len(sl5)-1])
	// 最初と最後の要素以外すべて
	fmt.Println(sl5[1 : len(sl5)-1])

	// 要素の追加
	sl6 := []string{"JAPAN", "USA"}
	new := append(sl6, "CHINA")
	fmt.Println(new)
	fmt.Println(sl6) // 元の値は変化なし

}
