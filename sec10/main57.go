package main

import "fmt"

func Double(i *int) {
	*i = *i * 3
}

func Double2(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

// コピーしないで渡すことができる =>参照渡し

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p) // 実態(値)

	n = 200
	fmt.Println(n) //=> 200
	fmt.Println(&n)
	fmt.Println(p)
	fmt.Println(*p) //=> 200

	*p = 300
	fmt.Println(n) //=> 300
	fmt.Println(&n)
	fmt.Println(p)
	fmt.Println(*p) //=> 300

	// *p = 300
	// fmt.Println(n)
	// fmt.Println(&n)
	// n = 200
	// fmt.Println(*p)
	// pとnは同じ値を指している

	/*
		Double(&n)
		fmt.Println(n)
		fmt.Println(&n)

		Double(p)
		fmt.Println(*p)
		fmt.Println(p)

		// 参照型のデータ構造(スライス)はもともと参照渡しの機能を持っているので、
		// ポインタ型を使用する必要がない
		var sl []int = []int{1, 2, 3}

		Double2(sl)
		fmt.Println(sl)

	*/

}
