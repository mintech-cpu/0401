package main

import "fmt"

type Point struct {
	A string
	B int
	C float64
}

func Update(p *Point) {
	p.A = "Update"
	p.B = 100
	p.C = 3.14
}

func main() {
	// 構造体を参照渡ししたい時は、こちらを推奨
	// アドレス演算子を使った構造体の生成
	p := &Point{}
	Update(p)
	fmt.Println(p)

	p2 := new(Point)
	Update(p2)
	fmt.Println(p2)
}
