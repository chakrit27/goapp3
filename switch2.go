package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine")
		fallthrough
	case 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
		fmt.Println("Buy some food")
		fallthrough
	case 31:
		fmt.Println("Party tonight")
	default:
		fmt.Println("No Tnformation for that day")
	}
}