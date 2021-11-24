package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scanln(&word)

	word = fmt.Sprintf("%s ", word)
	fmt.Println(strings.Repeat(word, 3))
}
