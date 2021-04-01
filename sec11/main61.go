package main

import "fmt"

// 構造体の埋め込み

type Point struct {
	A string
	B int
	C float64
}

type BigPoint struct {
	Point
}

// BigPointの構造体に、Pointの構造体を格納(埋め込み)

func (p *Point) Set(i int) {
	p.B = i
}

func main() {
	bp := BigPoint{}
	fmt.Println(bp)
	bp.Point.A = "banana"
	bp.Point.B = 300
	bp.Point.C = 3.14
	fmt.Println(bp)

	fmt.Println(bp.Point.A) // PointのAにアクセスしている

	// bpに値を格納
	bp.A = "apple"
	fmt.Println(bp.A)

	bp2 := BigPoint{}
	fmt.Println(bp2)

	bp2.Point.Set(500)
	fmt.Println(bp2)

}
