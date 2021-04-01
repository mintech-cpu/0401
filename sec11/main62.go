package main

import "fmt"

// 構造体のコンストラクタ
// pythonのinitのような役割

type Point struct {
	A string
	B int
	C float64
}

func NewPoint(a string, b int, c float64) *Point {
	return &Point{a, b, c}

}

// New構造体名

func main() {
	p1 := Point{"A", 1, 1.1}
	p2 := NewPoint("B", 2, 2.2)
	fmt.Println(p1)
	fmt.Println(p2)

}
