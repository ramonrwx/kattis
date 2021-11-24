package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

func ReadLines() []int {
        var output []int
        var input int
        fmt.Scanln(&input)

        s := bufio.NewScanner(os.Stdin)
        for s.Scan() {
                number, _ := strconv.Atoi(s.Text())
                output = append(output, number)

                if input == len(output) {
                        break
                }
        }

        return output
}

func main() {
        numbers := ReadLines()

        for i := len(numbers); i > 0; i-- {
                fmt.Println(numbers[i-1])
        }
}