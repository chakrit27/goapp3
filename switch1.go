package main

import "fmt"

func main() {
	grade := ""
	message := ""
	var marks int
	fmt.Print("Please Enter Marks: ")
	fmt.Scan(&marks)
	switch marks {
	case 80:
		grade = "A"
	case 70:
		grade = "B"
	case 60:
		grade = "C"
	default:
		grade = "ERROR"
	}
	fmt.Print(grade)

	switch {
	case grade == "A":
		message = "Excellent!\n"
	case grade == "B":
		message = "Well done\n"
	}
	fmt.Println(":", message)
}