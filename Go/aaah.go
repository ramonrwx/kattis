package main

import (
        "fmt"
)

func main() {
        var inputA, inputB string
        fmt.Scanln(&inputA)
        fmt.Scanln(&inputB)

        if len(inputA) < len(inputB) {
                fmt.Println("no")
        } else {
                fmt.Println("go")
        }
}