package main

import "fmt"

type student struct {
	name  string
	grade int
	profile
}

type profile struct {
	rumah string
	age int
}

func main() {
	var s1 student
	s1.name = "Dipta"
	s1.grade = 2
	s1.profile.rumah = "Tamanwinagnun"
	s1.profile.age = 12

	var s2 = student{"Kyoshi", 3, profile{"Kebumen", 13}}

	fmt.Println("Name\t: ", s1.name)
	fmt.Println("Kelas\t: ", s1.grade)
	fmt.Println("Rumah\t: ", s1.rumah)
	fmt.Println("Umur\t: ", s1.age)
	
	// fmt.Println("Name\t: ", s2.name)
	// fmt.Println("Kelas\t: ", s2.grade)
	// fmt.Println("Rumah\t: ", s2.rumah)
	// fmt.Println("Umur\t: ", s2.age)

	var s3 = &s2
	s2.name = "bilqis"
	fmt.Println("Name\t: ", s3.name)
	fmt.Println("Kelas\t: ", s3.grade)

	var orangbaru = struct{
		student
	}{}

	orangbaru.name = "Radit"
	orangbaru.grade = 4
	orangbaru.rumah = "Jkt"
	orangbaru.age = 30

	fmt.Println("Coba ah ", orangbaru.name)

}