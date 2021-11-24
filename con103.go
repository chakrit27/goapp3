package main

import "fmt"

func main() {
	age := 15
	national := "Thailand"
	if national == "Thailand" {
		fmt.Println("Welcome to Electric vote")
		if age >= 18 {
			fmt.Println("Eligible to give vote")
		}else{
			fmt.Printf("Not Eligible to give vote: %v < 18", age)
		}
	}

}