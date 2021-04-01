package main

import "fmt"

// 構造体のメソッド

type Point struct {
	A string
	B int
	C float64
}

func Set(p *Point, s string) {
	p.A = s
}

// メソッド => 構造体Point専用のメソッド、他の構造体では使用できない
func (p *Point) Set(s string) {
	p.A = s
}

// メソッドのレシーバー型(p *Point)はポインタ型にする
// => 参照渡しが可能になる

// func (構造体の指定 ポインタ型)Set(引数 データ型)
func main() {
	p1 := &Point{A: "Hello"}
	fmt.Println(p1)
	// Aの書き換え
	Set(p1, "Go")
	fmt.Println(p1)

	// メソッド
	p1.Set("Yes")
	fmt.Println(p1)
}
