package main

import "fmt"

func main() {
	var x, y = 10, 15
	fmt.Printf("%v == %v: %v\n", x, y, x == y) //false
	fmt.Printf("%v == %v: %v\n", x, y, x != y) //true
	fmt.Printf("%v == %v: %v\n", x, y, x > y) //false
	fmt.Printf("%v == %v: %v\n", x, y, x >= y) //false
	fmt.Printf("%v == %v: %v\n", x, y, x < y) //true
	fmt.Printf("%v == %v: %v\n", x, y, x <= y) //true
}