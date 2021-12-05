package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {
        s := bufio.NewScanner(os.Stdin)
        s.Scan()
        input := s.Text()

        if input == "OCT 31" || input == "DEC 25" {
                fmt.Println("yup")
        } else {
                fmt.Println("nope")
        }
}