package main

import (
    "os"
    "bufio"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var sum int
    for scanner.Scan() {
        value, _ := strconv.Atoi(scanner.Text())
        sum += value
    }
    sumStr := strconv.Itoa(sum)
    os.Stdout.WriteString(sumStr)
}