package main

import "fmt"

// スライス
/*
固定長である配列と異なり、スライスは可変長であるため柔軟にデータ（要素）を格納することが可能
*/

func main() {
	// 配列
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	/* 実行結果が[1 2 3 4 5]となるスライスを作成してください */
	// スライスの書き方は三通り
	// ①明示的
	var sl []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sl)
	// ②暗示的
	/*sl2 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl2)
	// ③makeの使用
	sl3 := make([]int, 5)
	fmt.Println(sl3)
	sl3[0] = 1
	sl3[1] = 2
	sl3[2] = 3
	sl3[3] = 4
	sl3[4] = 5
	fmt.Println(sl3)
	*/

	// 要素の取得
	// slの最初の要素を取得してください
	fmt.Println(sl[0])
	// slの2番目から3番目の要素を取得してください
	fmt.Println(sl[1:3])
	// slの最後の要素を取得してください
	fmt.Println(sl[len(sl)-1])
	// slの最初と最後の要素以外すべてを取得してください
	fmt.Println(sl[1 : len(sl)-1])

	// slの要素数を数えてください
	fmt.Println(len(sl))
	// slのキャパシティ(容量)を調べてください
	fmt.Println(cap(sl))

	// slに要素[6 7]を追加してください
	fmt.Println(append(sl, 6, 7))

	/* sl3[apple banana orange]をコピーして、sl4を作成した後、sl4の最初の要素をpeachに変更してください */
	// ①sl3をコピーし、sl４を作成
	sl3 := []string{"apple", "banana", "orange"}
	sl4 := make([]string, 3)

	n := copy(sl4, sl3)

	// ②sl4の最初の要素をappleからpeachに変更
	sl4[0] = "peach"
	fmt.Println(sl3) //sl3からsl4をコピーしただけなので、sl3の値は変更されず
	fmt.Println(sl4) // コピー後の値
	fmt.Println(n)   // コピーの要素数

}
