package main

import "fmt"

func main() {
	var cases, total int

	fmt.Scanln(&cases)
	for i := 1; i <= cases; i++ {
		var length int
		fmt.Scanln(&length)

		if i == 1 {
			total = length
			continue
		}

		total += length - 1
	}

	fmt.Println(total)
}
