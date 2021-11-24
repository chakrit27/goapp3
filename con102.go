package main

import "fmt"

func main() {
	isAdmin := true
	canUpdate := true

	if isAdmin == true {
		fmt.Println("Welcome,Admin")
		if canUpdate == true {
			fmt.Println("Can Update")
		}
	}
}