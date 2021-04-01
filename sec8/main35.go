package main

import "fmt"

// for

/*
 */

func main() {

	/*
		i := 0
		for i <= 10 {
			fmt.Println(i)
			i++
		}
	*/

	/*
		fmt.Println("***************")

		for i2 := 0; i2 <= 10; i2++ {
			if i2 == 3 {
				continue // スキップ
			}
			if i2 == 5 {
				break // 終了
			}
			fmt.Println(i2)
		}
	*/

	/*
		// 配列
		arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < len(arr); i++ {
			if i == 2 {
				continue
			}
			if i == 9 {
				break
			}
			fmt.Println(arr[i])
		}


		arr2 := [3]int{1, 2, 3} // 一つ目の返り値にインデックス番号/二つ目の返り値に要素の値
		for i, v := range arr2 {
			fmt.Println(i, v)
		}
	*/

	// スライス
	sl := []string{"A", "B", "C"}
	for i2, v2 := range sl {
		fmt.Println(i2, v2)
	}
	// マップ
	m := map[string]int{"A": 100, "B": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
