package main

import "fmt"

func main(){
	var names[5]string
	names[0] = "Dipta"
	names[1] = "Har"
	names[2] = "im"
	names[3] = "bawa"
	names[4] = "ez"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Print("Jumlah ada ", len(names), " element\n")
	fmt.Print(names, "\n")

	var buah = [...]string{"Melon", "Semangka", "Anggur", "Blueberry", "Naga"}
	fmt.Print(buah, "\n")
	fmt.Print("Jumlah buah ada : ", len(buah))

	for i := 0; i < len(buah); i++ {
		fmt.Println("Ini Buah : ", buah[i])
	}

	for i, satuan := range buah {
		fmt.Printf("elemen %d : %s\n", i, satuan)
	} 

	for _, satuan := range buah {
		fmt.Printf("Element ini bervalue : %s\n", satuan)
	}

	var hewan = [...][3]string{{"Monyet", "Kerbau", "Sapi"}, {"Pari", "Lele", "Cumi"}}
	fmt.Print(hewan, "\n")

	var ular = make([]string, 2)
	ular[0] = "Sanca"
	ular[1] = "Cobra"
	fmt.Print(ular)

}