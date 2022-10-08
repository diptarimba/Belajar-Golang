package main

import "fmt"

func main(){

	// Belajar Menggunakan If Else
	var angka = 4900.4

	if persen := angka/100; persen > 70 {
		fmt.Print("Golongan Angka Tinggi ", persen)
	} else if persen > 60 {
		fmt.Printf("Nilainya saat ini %.3f", persen)
	} else {
		fmt.Println("Nilainya hanya ", persen)
	}

	// Belajar Menggunakan Switch Case
	var bilangan = 9

	switch bilangan {
	case 7:
		fmt.Println("Benar Bilangan Ganjil")
	case 1,2,3,4,5,6,8:
		fmt.Println("Aneh Ternyata")
	default:
		{
			fmt.Println("9 Ternyata")
			fmt.Println("Bingung Ya")
		}
	}

	/*

	Belajar Menggunakan switch case
	model if-else

	*/

	angka = 8

	switch {
	case angka == (2 * 4):
		{
			fmt.Println("Angkanya benar benar :", angka )
			fmt.Println("Brilliant!")
			if bijian := angka; bijian > 6{
				hasilnya := bijian / 2
				fmt.Println("Hasilnya adalah", hasilnya)
			}
		}
	case angka < 4:
		{
			switch{
			case angka / 3 == 1:
				{
					fmt.Println("Ternyata Angka tsb adalah 3")
					fmt.Println("Keren Dips")
				}
			case angka == 2:
				fmt.Println("Ternyata cuma angka 2")
			default:
				fmt.Println("Angka 1 Dong...")
			}
		}
	}
}