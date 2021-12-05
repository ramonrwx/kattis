package main

import (
        "bufio"
        "fmt"
        "os"
)

func readLines() []string {
        var words []string
        var inputA int
        fmt.Scanln(&inputA)

        s := bufio.NewScanner(os.Stdin)

        for s.Scan() {
                words = append(words, s.Text())
                if inputA == len(words) {
                        break
                }
        }

        return words
}

func main() {
        words := readLines()

        for k, v := range words {
                if k%2 == 0 {
                        fmt.Println(v)
                }
        }
}