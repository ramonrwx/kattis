package main

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
)

func ReadLines() (phrases []string) {
        var cases int
        fmt.Scanln(&cases)

        s := bufio.NewScanner(os.Stdin)
        for s.Scan() {
                phrases = append(phrases, s.Text())
                if cases == len(phrases) {
                        break
                }
        }
        return
}

func main() {
        phrases := ReadLines()
        re := regexp.MustCompile(`^Simon says`)

        for _, phrase := range phrases {
                if re.MatchString(phrase) {
                        idx := re.FindStringIndex(phrase)[1] + 1
                        fmt.Println(phrase[idx:])
                }
        }
}