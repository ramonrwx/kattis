package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func readLines() (int, []string) {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        inputA, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        inputB := strings.Split(scanner.Text(), " ")

        return inputA, inputB
}

func convert(inputB []string) []int {
        output := make([]int, len(inputB))

        for i, value := range inputB {
                output[i], _ = strconv.Atoi(value)
        }
        return output
}

func main() {
        inputA, inputB := readLines()

        if inputA != len(inputB) {
                panic("inputA != inputB")
        }

        outputSum := 0

        for _, value := range convert(inputB) {
                outputSum += value
        }

        fmt.Println(outputSum)
}