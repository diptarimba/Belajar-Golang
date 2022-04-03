package main

import (
	"fmt"
	"strings"
)

type FilterCallback func(string) bool

func main() {
	var data = []string{"dipta", "jono", "kuni", "kono", "kano", "paino"}

	var datane = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var theteg = filter(data, func(each string) bool{
		return len(each) == 4
	})

	fmt.Println("Data mengandung huruf o :", datane)
	fmt.Printf("Yang hurufnya 4 adalah %v", theteg)
}

// func filter(data []string, callback func(string) bool) []string {
func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}