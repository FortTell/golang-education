package main

import (
    "fmt"
    "strconv"
)

func main(){
    var n int
    fmt.Scan(&n)
    s := strconv.Itoa(n)
    digits := make(map[rune]bool)
    for _, char := range s {
        digits[char] = true
    }
    if len(digits) < len(s) {
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
}