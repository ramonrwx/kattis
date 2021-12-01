package main

import "fmt"

func main() {
	var number string
	fmt.Scanln(&number)

	if number[:3] == "555" {
		fmt.Println(1)
		return
	}
	fmt.Println(0)
}
