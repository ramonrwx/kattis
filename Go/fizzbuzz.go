package main

import (
        "fmt"
)

func main() {
        var inputA, inputB, inputC int
        fmt.Scanf("%d %d %d", &inputA, &inputB, &inputC)

        for i := 1; i <= inputC; i++ {
                if i%inputA == 0 && i%inputB == 0 {
                        fmt.Println("FizzBuzz")
                } else if i%inputB == 0 {
                        fmt.Println("Buzz")
                } else if i%inputA == 0 {
                        fmt.Println("Fizz")
                } else {
                        fmt.Println(i)
                }
        }
}