package main

import "fmt"

func Double(i *int) {
	*i = *i * 5
}

func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)
	fmt.Println(p)
	fmt.Println(*p)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	/*
		n = 200
		fmt.Println(n)
		fmt.Println(&n)
		fmt.Println(p)
		fmt.Println(*p)

		*p = 300
		fmt.Println(n)
		fmt.Println(&n)
		fmt.Println(p)
		fmt.Println(*p)
	*/

	Double(&n)
	fmt.Println(n)

}
