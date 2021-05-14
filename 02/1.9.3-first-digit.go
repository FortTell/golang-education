package main

import (
    "fmt"
)

func main(){
    var n int
    fmt.Scan(&n)
    switch {
        case n < 10: fmt.Println(n)
        case n < 100: fmt.Println(n / 10)
        case n < 1000: fmt.Println(n / 100)
        default: fmt.Println(n / 1000)
    }
}