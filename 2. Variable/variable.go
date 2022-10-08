package main

import "fmt"

func main() {
	var anjay string = "Dipta"

	var kok string
	kok = "18"

	keterangan := "cukup tua"

	fmt.Printf("halo saya %s berumur %s, ya sudah %s!\n", anjay, kok, keterangan)

	var varpertama, varkedua, varketiga string

	varpertama = "kadal"

	varkedua = "gurun"

	varketiga = "makan"

	fmt.Printf("saya lihat %s %s lagi %s kaktus\n", varpertama, varkedua, varketiga)

	inibaru, inilagi, inilah := "anjay", 12, 2.2

	fmt.Printf("mencoba %s untuk melihat angka %d ternyata ketika dekat %f\n", inibaru, inilagi, inilah)
	fmt.Println("mencoba", inibaru, "untuk melihat angka", inilagi, "ternyata ketika dekat", inilah)

	_ = "nganu gan"

	percobaan, _ := "bersama", "saya"

	fmt.Printf("jalan %s saya\n", percobaan)

	lahhh := new(string)

	fmt.Println(lahhh)
	fmt.Println(*lahhh)
}