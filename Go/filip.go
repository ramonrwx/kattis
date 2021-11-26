package main

import (
        "fmt"
)

func reverse(num int) int {
        reversed := 0

        for num != 0 {
                reversed = reversed * 10
                reversed = reversed + num%10
                num = (num / 10)
        }

        return reversed
}

func main() {
        var inputA, inputB int
        fmt.Scanf("%d %d", &inputA, &inputB)

        inputA = reverse(inputA)
        inputB = reverse(inputB)

        if inputA > inputB {
                fmt.Println(inputA)
        } else {
                fmt.Println(inputB)
        }
}