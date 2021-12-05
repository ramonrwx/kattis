package main

import (
        "fmt"
)

func min(t, c, g int) (smallNum int) {
        if t <= c && t <= g {
                smallNum = t
        } else if c <= g {
                smallNum = c
        } else {
                smallNum = g
        }
        return
}

func main() {
        var input string
        var t, c, g int
        fmt.Scanln(&input)

        for i := 0; i < len(input); i++ {
                letter := string(input[i])

                if letter == "T" {
                        t++
                } else if letter == "C" {
                        c++
                } else if letter == "G" {
                        g++
                }
        }

        total := ((t * t) + (c * c) + (g * g)) + (7 * min(t, c, g))
        fmt.Println(total)
}