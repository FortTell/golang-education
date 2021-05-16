package main

import (
    "fmt"
)

func main() {
    var n int
    max := 0
    maxCount := 0
    for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
        if n == max {
            maxCount++
        } else if n > max {
            max = n
            maxCount = 1
        }
    }
    fmt.Println(maxCount)
}

