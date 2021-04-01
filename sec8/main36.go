package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 15:
		fmt.Println("young")
	case 30:
		fmt.Println("adult")
	default:
		fmt.Println("?")
	}

	age := 40
	switch {
	case age >= 0 && age < 3:
		fmt.Println("baby")
	case age >= 3 && age < 20:
		fmt.Println("young")
	case age >= 20 && age < 50:
		fmt.Println("adult")
	default:
		fmt.Println("old")
	}

	// 列挙型
	switch n2 := 15; n2 {
	case 15:
		fmt.Println("young")
	case 30:
		fmt.Println("adult")
	default:
		fmt.Println("?")
	}

	// 演算子
	m := 6
	switch {
	case m > 0 && m < 4:
		fmt.Println("0 < m < 4")
	case m > 3 && m < 7:
		fmt.Println("3 < v < 7")
	}
	// このように書く場合は、switchに与えるしきを省略
	// 列挙型と演算子を使用した式は混在できない => エラー

}
