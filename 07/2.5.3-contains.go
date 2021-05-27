package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    x := ReadAndTrimNewline(reader)
    s := ReadAndTrimNewline(reader)
    
    fmt.Println(strings.Index(x, s))
}