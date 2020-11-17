package main

import (
	"fmt"
	"reflect"
)

func main() {
	var item int

	if reflect.ValueOf(item).Type().Kind() != reflect.Ptr {
		fmt.Println("is not pointer")
	}
}
