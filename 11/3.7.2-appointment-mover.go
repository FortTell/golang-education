package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strings"
)

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    text = strings.TrimRight(text, "\r\n")
    layout := "2006-01-02 15:04:05"
    parsed, _ := time.Parse(layout, text)
    if parsed.Hour() >= 13 {
        parsed = parsed.Add(time.Hour * 24)
    }
    fmt.Println(parsed.Format(layout))
}