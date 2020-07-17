package main

import "fmt"

func main() {
	row := 3
	// col := 3

	a := make([][]int, row)
	// for i := range a {
	// 	a[i] = make([]int, col)
	// }

	fmt.Println(a)

	tryOther()
	tryMoreOfTheOther()
}

func tryOther() {
	x := [][]string{{}, {}, {}}

	fmt.Println(x)
}

func tryMoreOfTheOther() {
	x := [][]interface{}{{}, {}, {}}

	fmt.Println(x)

	fmt.Println([][]string{{}, {}, {}})
}
