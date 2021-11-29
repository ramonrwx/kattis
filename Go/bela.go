package main

import "fmt"

func main() {
	var c int
	fmt.Scanln(&c)
	c <<= 2
	fmt.Println(c)
}
