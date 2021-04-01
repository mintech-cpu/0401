package main

import (
	"fmt"
)

// 構造体とスライス
// Goで頻出

type Person struct {
	Name string
}

type Persons struct {
	Persons []*Person
}

func main() {
	/*
		ps := make([]Person, 5)
		fmt.Println(ps)

		ps[0].Name = "Mike"
		ps[1].Name = "Mike"
		ps[2].Name = "Mike"
		ps[3].Name = "Mike"
		ps[4].Name = "Mike"
		fmt.Println(ps)
	*/

	p1 := Person{"Mike"}
	p2 := Person{"Mike"}
	p3 := Person{"Mike"}
	p4 := Person{"Mike"}
	p5 := Person{"Mike"}

	ps := Persons{}
	ps.Persons = append(ps.Persons, &p1, &p2, &p3, &p4, &p5)

	for _, p := range ps.Persons {
		fmt.Println(p)
	}

	// スライスでデータを管理

}
