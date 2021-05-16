package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Scan(&c)
    fmt.Scan(&d)
    for fmt.Scan(&n); n > 100; fmt.Scan(&n) {
        if n >= 10 {
            fmt.Println(n)
        }
    }
}

