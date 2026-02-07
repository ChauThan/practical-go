package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	var i any

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = true
	describe(i)

	i = 3.14
	describe(i)
}