package main

import "fmt"

func main() {
	var letters string
	fmt.Scanln(&letters)

	ball := 1

	for _, letter := range letters {
		swap(letter, &ball)
	}

	fmt.Println(ball)
}

func swap(letter rune, ball *int) {
	if letter == 65 {
		if *ball == 1 {
			*ball = 2
		} else if *ball == 2 {
			*ball = 1
		}
	} else if letter == 66 {
		if *ball == 2 {
			*ball = 3
		} else if *ball == 3 {
			*ball = 2
		}
	} else {
		if *ball == 1 {
			*ball = 3
		} else if *ball == 3 {
			*ball = 1
		}
	}
}
