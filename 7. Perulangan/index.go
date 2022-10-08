package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		fmt.Println("Angka ", i)
	}

	var j = 0
	
	for {
		fmt.Println("Angkanya Sekarang : ", j)
		
		if j > 10 {
			break
		}
		j++
	}

	var k = 0

	for {
		if k % 2 == 0 {
			fmt.Println("Belajar memahami, ini adalah bilangan genap")
		} else {
			fmt.Println("Ini Angka : ", k)
		}
		

		if k == 10{
			break
		}
		k++
	}

	for l := 0; l < 5; l++{
		for m := l; m < 10; m++ {
			if m % 2 == 0 {
				continue
			}
			fmt.Print(m, " ")
		}
		fmt.Println("")
	}

	berhenti:
	for n := 0; n < 5; n++ {
		for o := 0; o < 5; o++ {
			for p := 0; p < 5; p++ {
				fmt.Print("Matriks [", n, "][", o, "][", p, "]\n")
				if n  == 4 && o  == 4 && p == 4 {
					break berhenti
				}
			}
		}
	}
}