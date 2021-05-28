package main

import (
    "fmt"
    "unicode"
)

func main() {
    text := ReadAndTrimSingleNewline()
    if len(text) < 2 {
        fmt.Println("Wrong")
        return
    }
    
    runes := []rune(text)
    if unicode.IsUpper(runes[0]) && runes[len(runes) - 1] == '.' {
        fmt.Println("Right")
    } else {
        fmt.Println("Wrong")
    }
}