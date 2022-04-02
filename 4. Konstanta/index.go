package main

import "fmt"

func main(){
	const nama string = "Dipta Harimbawa"

	fmt.Print("Nama Saya ", nama + "\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}