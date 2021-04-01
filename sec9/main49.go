package main

import "fmt"

func main() {

	// mapの作成方法は3通り
	// 明示関数
	var m = map[string]int{"A": 100, "B": 200, "C": 300}
	fmt.Println(m)

	// 暗示関数
	m2 := map[string]int{"A": 100, "B": 200, "C": 300}
	fmt.Println(m2)

	// 改行も可能
	m3 := map[int]string{
		1: "北海道",
		2: "青森県",
		3: "岩手県",
	}
	fmt.Println(m3)

	// make
	m4 := make(map[int]string)
	fmt.Println(m4)
	// 要素の追加
	m4[0] = "JAPAN"
	m4[1] = "USA"
	m4[3] = "CHINA"
	fmt.Println(m4)

	m5 := make(map[string]int)
	m5["A"] = 100
	m5["B"] = 200
	m5["C"] = 300
	fmt.Println(m5)

	// 要素の取得
	fmt.Println(m["A"])
	fmt.Println(m4[1])

	// 要素の削除
	delete(m4, 3) // 第一引数に削除したいmap, 第二引数にkey
	fmt.Println(m4)

	// 要素数を調べる
	fmt.Println(len(m4))

	// 値の取得(エラーハンドリング)
	// 値があるので値が返ってくる
	s := m4[1]
	fmt.Println(s)

	//　値がないので何も返ってこない
	s1 := m4[3]
	fmt.Println(s1)

	// s, ok := m4[3]
	// if !ok {
	// 	fmt.Println("error")
	// }
	// fmt.Println(s, ok)

	_, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok)

}
