package main

// switch型スイッチ

import "fmt"

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "?")
	}
}

func main() {
	anything("aaa")
	anything(1)

	// 型アサーション
	var x interface{} = 1.3

	/*
		i := x.(int) // 変数の.復元したい型
		fmt.Println(i)
		fmt.Printf("%T", i)
		fmt.Println(i + 3)
	*/

	i2, isInt := x.(int)
	fmt.Println(i2, isInt)

	// 二つの返り値を使用した型アサーションなら、異なる型で復元しても、エラーにならない
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("?")
	}

	// ifよりもswitchを推奨
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("?")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "?")
	}

}
