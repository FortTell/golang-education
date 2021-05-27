package main

import (
    "fmt"
)

func main() {
    s := ReadAndTrimSingleNewline()
    
    oddRunes := make([]rune, 0, len(s) / 2)
    for i, c := range s {
        if i % 2 != 0 {
            oddRunes = append(oddRunes, c)
        }
    }
    
    fmt.Println(string(oddRunes))
}