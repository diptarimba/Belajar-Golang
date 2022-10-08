package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func (s student) changeName1(name string) {
	fmt.Println("---> Nama", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> Nama", name)
	s.name = name
}

func main() {
	var s1 = student{"Dipta Harimbawa", 11}
	fmt.Println("Nama Asli", s1.name)

	s1.changeName1("Dipta Kawe")
	fmt.Println("Nama Ganti", s1.name)

	s1.changeName2("Dipta Gali")
	fmt.Println("Nama Lagi", s1.name)
}
