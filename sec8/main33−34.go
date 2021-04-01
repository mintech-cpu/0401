package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	if a == 1 {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

	if a := "apple"; a == "apple" {
		fmt.Println("100yen")
	} else if a == "banana" {
		fmt.Println("200yen")
	} else {
		fmt.Println("free")
	}

	// エラーハンドリング
	var s string = "100" // 代入の値がAだとエラー
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("%T", i)

}
