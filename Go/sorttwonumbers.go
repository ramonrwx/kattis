package main

import "fmt"

func main() {
        var inputA, inputB int
        fmt.Scanf("%d   %d", &inputA, &inputB)

        if inputA > inputB {
                fmt.Println(inputB, inputA)
        } else {
                fmt.Println(inputA, inputB)
        }
}