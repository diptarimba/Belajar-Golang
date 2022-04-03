package main

import (
	"fmt"
	"strings"
)

func main() {

	var rata2 = calculate(10,7,5,2,5,7,3,6,5,3,4,5,6,2,4,5,6,4,3,21,5,7,4,3,4,2,3)
	var hasil = fmt.Sprintf("1. Nilai rata-rata angka diatas adalah : %.2f", rata2)
	fmt.Println(hasil)

	var cow = []int{2,3,5,6,7,2,3,4,6,2,3,4,6,7,2,3,4,6,7,8,4,5,7,8,9,7,8,9,6,7}
	rata2 = calculate(cow ...)
	hasil = fmt.Sprintf("2. Nilai rata-rata angka diatas adalah : %.2f", rata2)
	fmt.Println(hasil)

	var katarumpang = []string{"dipta", "sangat", "ganteng"}
	var pergantengan = collate("yang namanya", katarumpang ...)
	fmt.Println(pergantengan)
}

func calculate(number ... int) (avg float64) {
	var total int = 0
	for _,i := range number {
		total += i
	}
	avg = float64(total) / float64(len(number))
	return
}

func collate(word string, multipleword ... string) (ret string){
	var gabungan string = strings.Join(multipleword, " ")
	ret = fmt.Sprint(word, " ", gabungan)
	return
}