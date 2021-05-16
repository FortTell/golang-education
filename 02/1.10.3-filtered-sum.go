package main

import (
    "fmt"
)

func main() {
    sum := 0
    var n, temp int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&temp)
        if temp % 8 == 0 && temp >= 10 && temp <= 99 {
            sum += temp
        }
    }
    fmt.Println(sum)
}

