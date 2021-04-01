package main

import "fmt"

func inter(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "?")
	}
}

func main() {
	// forループで0~9を表示してください
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 3をスキップ、6になったらforループを終了してください
	for i2 := 0; i2 < 10; i2++ {
		if i2 == 3 {
			continue
		} else if i2 == 6 {
			break
		} else {
			fmt.Println(i2)
		}
	}

	// 配列「arr := [5]int{1, 2, 3, 4, 5}」をforループで表示してください
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr2 := [3]string{"A", "B", "C"}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	sl := []string{"A", "B", "C"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	m := map[string]int{"A": 10, "B": 20}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 0~19才をyoung、20~65才未満をadult、65才以上をseniorとするswitch文を作成し、30才はどれに当てはまるか表示してください。
	age := 30
	switch {
	case age >= 0 && age < 20:
		fmt.Println("young")
	case age >= 20 && age < 65:
		fmt.Printf("adult")
	default:
		fmt.Println("senior")
	}

	// interface型の変数xを作成して、値を100としてください。
	var x interface{} = 100
	fmt.Println(x)

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("?")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println(v, "?")
	}

	inter("A")
	inter(100)
	inter(5.5)
}
