package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main(){

	rand.Seed(time.Now().Unix())
	var umursaya int

	var names = []string{"dipta", "harimbawa"}
	printmassage("Hello", names)

	umursaya = kirakira(10,22)
	fmt.Println("Umura saya adalah ", umursaya)

	prediksi(1,5,4)
}

func printmassage(message string, arr []string){
	fmt.Println(message, strings.Join(arr, " "))
}

func kirakira(batas1 int, batas2 int) int{
	var value = rand.Int() % (batas2 - batas1 + 1) + batas1
	return value
}

func prediksi(hitung, hitang, hiting int){
	for hitung := 0; hitung < hitang; hitung++ {
		fmt.Println("Ini Percobaan ke ", hitung)
		if hitung == hiting {
			return
		}
	}
}