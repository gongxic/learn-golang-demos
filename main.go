package main

import (
	"fmt"
	"reflect"
)

func main() {

	b := 12

	fmt.Println(b, 123, 123, 123, 123)
	typeofb := reflect.TypeOf(b)

	fmt.Println(typeofb)

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	b := 123
	fmt.Print(b)
}
