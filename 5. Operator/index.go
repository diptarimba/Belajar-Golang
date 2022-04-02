package main

import "fmt"

func main(){
	var value = (2 * 4) / (8 - (3 + 1))
	var isValue = (value == 2)

	fmt.Print("Nilainya kesamaannya adalah ", isValue, "\n")

	fmt.Printf("Nilai kesamaannya adalah (%t)", isValue)
}