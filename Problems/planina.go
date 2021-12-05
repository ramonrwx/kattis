package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	n = (1 << n) + 1
	fmt.Println(n * n)
}
