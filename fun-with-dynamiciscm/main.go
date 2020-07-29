package main

import "fmt"

func main() {
	row := 3
	// col := 3

	a := make([][]string, row)
	b := [][]string{
		{"testtest2"},
		{"test", "tttt"},
		{"test", "tttt", "teaet"},
	}
	// for i := range a {
	// 	a[i] = make([]int, col)
	// }

	fmt.Println(a)
	fmt.Println(b)

	// tryOther()
	// tryMoreOfTheOther()
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
