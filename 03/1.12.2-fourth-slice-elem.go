package main

import (
    "fmt"
)

func main() {
    var n, temp int
    fmt.Scan(&n)
    slice := make([]int, 4)
    for i := 0; i < n; i++ {
        fmt.Scan(&temp)
        if i < 4 {
            slice[i] = temp
        }
    }
    fmt.Println(slice[3])
}