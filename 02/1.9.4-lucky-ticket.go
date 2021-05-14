package main

import (
    "fmt"
    "strconv"
)

func main(){
    var n int
    fmt.Scan(&n)
    first := n / 1000
    last := n % 1000
    if sumOfDigits(first) == sumOfDigits(   last) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

func sumOfDigits(number int) int {
    sum := 0
    for _, digit := range strconv.Itoa(number) {
        sum += int(digit - '0')
    }
    return sum
}