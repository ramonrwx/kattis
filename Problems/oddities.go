package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
)

func ReadLines() []int {
        var input int
        var nums []int

        fmt.Scanln(&input)

        s := bufio.NewScanner(os.Stdin)
        for s.Scan() {
                num, _ := strconv.Atoi(s.Text())
                nums = append(nums, num)
                if input == len(nums) {
                        break
                }
        }

        return nums
}

func main() {
        nums := ReadLines()

        for _, v := range nums {
                if v%2 == 0 {
                        fmt.Println(v, "is even")
                } else {
                        fmt.Println(v, "is odd")
                }
        }
}