package main

import (
        "fmt"
        "strings"
)

func main() {
        var input string
        fmt.Scanln(&input)

        eNeeded := (len(input) - 2) * 2
        fmt.Printf("h%vy", strings.Repeat("e", eNeeded))
}