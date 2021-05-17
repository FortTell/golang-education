package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    var a, positiveCount int
    for i := 0; i < n; i++ {
        fmt.Scan(&a)
        if a > 0 {
            positiveCount++
        }
    }
    fmt.Println(positiveCount)
}