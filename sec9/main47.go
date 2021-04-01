package main

import "fmt"

// スライスのfor

func main() {

	sl := []string{"a", "b", "c"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	// 要素の取り出し方①
	for _, v := range sl {
		fmt.Println(v)
	}

	// 要素の取り出し方②
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}
