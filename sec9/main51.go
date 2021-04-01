package main

import "fmt"

// チャネル

func main() {
	// 双方向
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch3 chan<- int

	// チャネルの初期化
	ch1 = make(chan int)
	ch2 := make(chan int)
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	// バッファサイズ(チャネルの容量)を設定
	ch3 := make(chan int, 5) // バッファサイズ(この場合は5)を超えたデータを送ると、deadlock
	fmt.Println(cap(ch3))

	// チャネルにデータを送信(追加)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println(len(ch3)) // len()でデータの要素数を調べる

	// チャネルからデータを受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("チャネルの要素数", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("チャネルの要素数", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("チャネルの要素数", len(ch3))

}
