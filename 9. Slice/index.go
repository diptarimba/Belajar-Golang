package main

import "fmt"

func main() {
	var fruits = [][]string{{"apple", "anggur"}, {"iguana", "sapi"}}
	fmt.Println(fruits);

	var fruit = []string{"apple", "grape", "banana", "melon"}
	var freed = append(fruit, "sawi")
	var newFruits = fruit[0:2]
	fmt.Println(newFruits)
	fmt.Println(freed)

	fmt.Println(fruit[0:2])
	fmt.Println(fruit[0:4])
	fmt.Println(fruit[0:1])
	fmt.Println(fruit[2:4])

	var freedbaru = make([]string,3)
	var _ = copy(freedbaru, freed)
	var nanatsu = []string{"kadal", "kelinci"}
	var freedbaru2 = copy(freedbaru, nanatsu)

	fmt.Println(freedbaru)
	fmt.Println(freedbaru2)

	fmt.Println(len(freedbaru)) // panjang array
	fmt.Println(cap(freedbaru)) // kapasitas sebenarnya
}