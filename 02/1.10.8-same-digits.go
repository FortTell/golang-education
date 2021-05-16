package main

import (
    "fmt"
    "strconv"
)

func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    bDigits := make(map[rune]bool)
    for _, char := range strconv.Itoa(b) {
        bDigits[char] = true
    }
    for _, char := range strconv.Itoa(a) {
        if bDigits[char] {
            fmt.Print(char - '0')
            fmt.Print(" ")
        }
    }
}

