package main

import (
    "fmt"
)

func work(x int) int {
    //sleep
    return x * 2
}

func main() {
    cache := make(map[int]int)
    for i := 0; i < 10; i++ {
        var curr int
        fmt.Scan(&curr)
        value, exists := cache[curr]
        if !exists {
            value = work(curr)
            cache[curr] = value
        }
        fmt.Printf("%d ", value)
    }
}