package main

import "fmt"

func main() {
	var x = 10
	var y int = 20

	if true {
		fmt.Println("True")
	}

	if x < y {
		fmt.Println("X is less than y")
	}
}