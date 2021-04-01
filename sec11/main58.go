package main

import "fmt"

// 構造体

type Point struct {
	A int
	B string
	C float64
}

func Update(p Point) {
	p.A = 100
	p.B = "Update"
	p.C = 3.14
}

func main() {

	var p4 Point
	Update(p4)
	fmt.Println(p4)

	// 明示的(初期値)
	var p Point
	fmt.Println(p)

	p2 := Point{A: 1, B: "Hello", C: 2.2}
	fmt.Println(p2)
	fmt.Println(p2.A)

	// フィールドの値を上書き
	p2.A = 100
	fmt.Println(p2.A)
	fmt.Println(p2)

	// 構造体のフィールドの順番
	p3 := Point{22, "yes", 5.5} // {}の中は、構造体のフィールドの定義の順番に宣言する必要がある
	fmt.Println(p3.A)

}
