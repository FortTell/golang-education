package main

import (
    "fmt"
    "time"
)

func main() {
    var dateStr string
    fmt.Scan(&dateStr)
    parsed, _ := time.Parse(time.RFC3339, dateStr)
    fmt.Println(parsed.Format(time.UnixDate))
}