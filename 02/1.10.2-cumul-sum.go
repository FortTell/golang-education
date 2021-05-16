package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(sumUntilNumber(b) - sumUntilNumber(a))
}

func sumUntilNumber(number int) int {
    return (number * (number + 1)) / 2
}
