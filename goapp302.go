package main

import "fmt"

func main() {
	var x int = 200
	var y int = -1000
	var m uint = 2000
	var n float32 = 123.52

	var txt1 string = "Hello"

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	fmt.Printf("Type: %T, value: %v\n", m, m)
	fmt.Printf("Type: %T, value: %v\n", n, n)
	fmt.Printf("Type: %T, value: %s\n", txt1, txt1)
}