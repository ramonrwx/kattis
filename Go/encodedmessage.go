package main

import (
        "bufio"
        "fmt"
        "math"
        "os"
)

func main() {
        var (
                cases int
                words []string
        )

        fmt.Scanln(&cases)

        s := bufio.NewScanner(os.Stdin)
        for s.Scan() {
                words = append(words, s.Text())
                if cases == len(words) {
                        break
                }
        }

        for i := 0; i < cases; i++ {
                length := int(math.Sqrt(float64(len(words[i]))))
                var output string

                for x := length - 1; x >= 0; x-- {
                        for y := x; y < len(words[i]); y += length {
                                output += string(words[i][y])
                        }
                }

                fmt.Println(output)
        }
}