package main

import "fmt"

func main(){
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Nilai A\t:", numberA)
	fmt.Println("Nilai A\t:", &numberA)

	fmt.Println("Nilai B\t:", *numberB)
	fmt.Println("Nilai B\t:", numberB)
}