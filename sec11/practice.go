package main

import "fmt"

type Book struct {
	Type  string
	Price int
}

func Update(p *Book) {
	p.Type = "comics"
	p.Price = 550
}

func (p *Book) Set(s string, i int) {
	p.Type = s
	p.Price = i
}

func main() {
	var p Book
	p.Type = "magazine"
	p.Price = 500
	fmt.Println(p)

	p2 := Book{Type: "magazine", Price: 500}
	fmt.Println(p2)

	p3 := Book{"magazine", 500}
	fmt.Println(p3)

	// pの値段を取得
	fmt.Println(p.Price)

	p4 := &Book{}
	Update(p4)
	fmt.Println(p4)

	p5 := &Book{Type: "magazine", Price: 550}
	fmt.Println(p5)
	p5.Set("novel", 300)
	fmt.Println(p5)
}
