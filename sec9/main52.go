package main

import "fmt"

// チャネルと後ルーチン
func main() {
	ch1 := make(chan int)
	fmt.Println(<-ch1)
}
