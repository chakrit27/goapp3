package main

import "fmt"

func main() {
	var x, y, a, b = 40, 10, 20, 30

	fmt.Printf("x + y = %d\n", x + y)
	fmt.Printf("x - y = %d\n", x - y)
	fmt.Printf("x * y = %d\n", x * y)
	fmt.Printf("x / y = %d\n", x / y)
	fmt.Printf("x mod y = %d\n", x % y)

	//post
	a++ //a = 20
	b-- //a = 30
	fmt.Printf("a++ = %d\n", a) //21 = 20 + 1
	fmt.Printf("b-- = %d\n", b) //29 = 30 - 1
}