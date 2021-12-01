package main

import "fmt"

func main() {
	var n, r int = 0, 0

	fmt.Scanln(&n)
	for n != 0 {
		r = (r << 1) | (n & 0x1)
		n = n >> 1
	}

	fmt.Print(r)
}
