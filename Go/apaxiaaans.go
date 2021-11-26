package main

import (
        "bytes"
        "fmt"
)

func main() {
        var (
                input  string
                last   rune
                output bytes.Buffer
        )

        fmt.Scanln(&input)

        for i, letter := range input {
                if letter != last || i == 0 {
                        output.WriteRune(letter)
                        last = letter
                }
        }

        fmt.Println(output.String())
}