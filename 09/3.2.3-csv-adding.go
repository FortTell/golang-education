package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
    "encoding/csv"
    "strings"
    "strconv"
)

func prepareNumber(field string) float64 {
    cleanedField := strings.ReplaceAll(strings.ReplaceAll(field, " ", ""), ",", ".")
    number, err := strconv.ParseFloat(cleanedField, 64)
    if err != nil {
        panic(err)
    }
    return number
}

func main() {
    csv := csv.NewReader(bufio.NewReader(os.Stdin))
    csv.Comma = ';'
    for {
        record, err := csv.Read()
        if err == io.EOF {
            break
        }
        dividend := prepareNumber(record[0])
        divider := prepareNumber(record[1])
        fmt.Printf("%.4f", dividend / divider)
    }
}