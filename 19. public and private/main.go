package main

import (
	"fmt"
	// "public-and-private/library"
	// . "public-and-private/library"
	l "public-and-private/library"
)

func main() {
	// library.SayHello("etdah")

	// var s1 = library.Student{"dino", 12}
	// var s1 = Student{"dino", 12}
	var s1 = l.Student{"dino", 12}
	fmt.Println(s1.Name)
	fmt.Println(s1.Grade)
}
