package main

import (
	"fmt"
	"math"
)

func main(){
	
	var luaskubus, volumekubus = calculate(2)
	fmt.Println("Luas Kubusnya adalah :", luaskubus)
	fmt.Println("Volume Kubusnya adalah :", volumekubus)

	var luasbola, volumebola = bolahitung(5)

	fmt.Println("Luas bolanya adalah ", luasbola)
	fmt.Println("Volume bolanya adalah ", volumebola)
}

func calculate(ruas float32)(float32,float32){
	var luasnya = math.Pow(2,float64(ruas)) * 6
	var volumnya = math.Pow(3,float64(ruas))

	return float32(luasnya),float32(volumnya)
}

func bolahitung(jarijari float64)(luasbola, volumebola float64){
	luasbola = 4 * math.Pi * math.Pow(2,jarijari)
	volumebola = 4/3 * math.Pi * math.Pow(3, jarijari)
	return
}