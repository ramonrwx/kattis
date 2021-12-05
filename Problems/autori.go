package main

import (
        "fmt"
        "strings"
)

func main() {
        var input string
        fmt.Scanln(&input)

        short := strings.Split(input, "-")
        for _, letter := range short {
                fmt.Print(letter[:1])
        }
}