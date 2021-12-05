package main

import (
        "fmt"
        "sort"
)

func main() {
        var a, b, c int
        var letter string

        fmt.Scanln(&a, &b, &c)
        fmt.Scanln(&letter)

        var nums = []int{a, b, c}
        sort.Ints(nums)

        for i := 0; i < 3; i++ {
                if string(letter[i]) == "A" {
                        fmt.Printf("%d ", nums[0])
                } else if string(letter[i]) == "B" {
                        fmt.Printf("%d ", nums[1])
                } else {
                        fmt.Printf("%d ", nums[2])
                }
        }
}