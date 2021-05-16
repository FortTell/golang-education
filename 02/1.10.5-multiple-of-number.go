package main

import (
    "fmt"
)

func main() {
    var n, c, d int
    fmt.Scan(&n)
    fmt.Scan(&c)
    fmt.Scan(&d)
    for i := 0; i <= c; i += c {
        if i % d != 0 {
            fmt.Println(i)
            break
        }
    }
}

