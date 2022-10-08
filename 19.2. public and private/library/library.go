package library

import "fmt"

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Sujono Soekanto"
	Student.Grade = 3

	fmt.Println("menghadehhhh library imported")
}
