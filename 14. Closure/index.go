package main

import "fmt"

func main() {

	var infoJurusan = func(n []int) (hasil []string) {
		for _, i := range n {
			switch i {
			case 1:
				hasil = append(hasil, "Jurusan Akuntansi")
			case 2:
				hasil = append(hasil, "Jurusan Mesin")
			case 3:
				hasil = append(hasil, "Jurusan Elektro")
			case 4:
				hasil = append(hasil, "Jurusan Sipil")
			default:
				hasil = append(hasil, "Jurusan Administrasi Bisnis")
			}
		}

		return
	}

	var urutan = []int{1, 5, 4, 3, 2}
	var urutkan = infoJurusan(urutan)
	for _, urut := range urutkan {
		fmt.Printf("%v\n", urut)
	}

	var nilaiuang = []int{1000,2000,3000,4000,5000}

	var panjang, ratauang = func(n int) (int, func() []int) {
		var kembalian []int
		for i, nilai := range nilaiuang {
			if(i < n){
				continue
			}
			kembalian = append(kembalian, nilai * 10)
		}
		return len(kembalian), func() []int {
			return kembalian
		}
	}(1)

	var angkanya = ratauang()

	fmt.Println("Perubahan ", angkanya)
	fmt.Println("Aslinya", nilaiuang)
	fmt.Println("Panjang", panjang)

}