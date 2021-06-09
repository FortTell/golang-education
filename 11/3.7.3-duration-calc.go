package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    dateTime1, _ := reader.ReadString(',')
    dateTime1 = strings.TrimRight(dateTime1, ",")
    dateTime2, _ := reader.ReadString('\n')
    dateTime2 = strings.TrimRight(dateTime2, "\r\n")
    
    layout := "02.01.2006 15:04:05"
    parsedDT1, _ := time.Parse(layout, dateTime1)
    parsedDT2, _ := time.Parse(layout, dateTime2)
    
    if parsedDT1.After(parsedDT2) {
        fmt.Println(parsedDT1.Sub(parsedDT2).Round(time.Second))
    } else {
        fmt.Println(parsedDT2.Sub(parsedDT1).Round(time.Second))
    }
}