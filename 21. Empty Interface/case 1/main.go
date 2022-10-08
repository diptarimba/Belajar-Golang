package main

import "fmt"

func main() {
	var secret interface{}

	secret = "wadhu wadhu"
	fmt.Println(secret)

	secret = []string{"pisag", "anggur", "jeruk"}
	fmt.Println(secret)

	secret = 13.4
	fmt.Println(secret)

	var data map[string]any

	data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
}
