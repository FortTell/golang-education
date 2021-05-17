package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    var a int
    for i := 0; i < n; i++ {
        fmt.Scan(&a)
        if i % 2 == 0 {
            fmt.Printf("%d ", a)
        }
    }
}