package main

import (
    "os"
    "bufio"
    "strings"
)

func ReadAndTrimNewline(reader *bufio.Reader) string {
    s, _ := reader.ReadString('\n')
    return strings.Trim(s, "\r\n")
}

func ReadAndTrimSingleNewline() string {
    s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    return strings.Trim(s, "\r\n")
}