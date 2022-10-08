package main

import "fmt"

func main(){
	var angkasaja map[string]int
	angkasaja = map[string]int {}
	angkasaja["pertama"] = 1
	angkasaja["kedua"] = 2
	angkasaja["ketiga"] = 3

	fmt.Println(angkasaja["kedua"])

	var namateman = map[string]string {
		 "nama" : "temanluffy",
		 "kelas" : "iyabisa",
	}

	fmt.Printf("Nama saya %s sekarang kelas %s\n", namateman["nama"], namateman["kelas"])

	//deklarasi map baru

	var cupu = map[string]string{}
	var cucu = make(map[string]string)
	var cuci = *new(map[string]string)

	fmt.Println(cupu,cucu,cuci)

	var nilaikelas = map[string]int {
		"dipta" : 10,
		"agvirlia" : 90,
		"disni" : 90,
		"ridho" : 89,
		"siapa" : 1,
	}

	for i,c := range nilaikelas {
		fmt.Println("Nilai ", i + "\t: ", c)
	}
	
	fmt.Println("Setelah penghapusan")
	delete(nilaikelas, "siapa")

	for i,c := range nilaikelas {
		fmt.Println("Nilai ", i + "\t: ", c)
	}

	var multidimensional = []map[string]string{
		// map[string]string{"nama" : "dipta", "kelas" : "am4a"}, // Baku
		// map[string]string{"nama" : "disni", "kelas" : "am4a"}, // Baku
		// map[string]string{"nama" : "anisa", "kelas" : "am4a"}, // Baku
		{"nama" : "dipta", "kelas" : "am4a"},
		{"nama" : "disni", "kelas" : "am4a"},
		{"nama" : "anisa", "kelas" : "am4a"},
	}

	for _,val := range multidimensional {
		fmt.Println("Nama : ", val["nama"], " Kelas : ", val["kelas"])
	}

	var multidimensionalraw = []map[string]string {
		{"nama" : "dipta"},
		{"alumni" : "smansa", "kelas" : "IIS B"},
	}

	for _,value := range multidimensionalraw {
		for key,values := range value {
			fmt.Println(key, "\t\t", values)
		}
		fmt.Print("\n")
	}


}