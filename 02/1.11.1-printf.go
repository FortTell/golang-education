package main

import (
    "fmt"
)

func main() {
    var r float64
    fmt.Scan(&r)
    switch {
        case r <= 0: fmt.Printf("число %2.2f не подходит", r)
        case r > 10000: fmt.Printf("%e", r)
        default: fmt.Printf("%.4f", r * r)
    }
}

