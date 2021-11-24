package main

import "fmt"

func main() {
	fmt.Print("Please Enter Countries as th,ma: ")
	province := ""
	fmt.Scan(&province)

	switch province {
	case "th":
		fmt.Println("Thailand")
	case "ma":
		fmt.Println("Malaysia")
	default:
		fmt.Println("ERROR")
	}
}