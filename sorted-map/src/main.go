package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    textFile := "../gwtw.txt"
    file, err := os.Open(textFile)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)
    //wordCountMap := make(map[string]int)
    //keysList := make([]string, 0)
    sm := SortedMap{wordCount: make(map[string]int), keysByInsertOrder: make([]string, 0)}
    for scanner.Scan() {
        line := strings.Trim(scanner.Text(), " \t")
        if len(line) == 0 {
            continue
        }
        for _, word := range strings.Split(line, " ") {
            prepared := strings.ToLower(strings.Trim(word, "'?!,.:"))
            sm.AddWord(prepared)
        }
    }
    sm.GetWords(10, 10)
}

type SortedMap struct {
    wordCount map[string]int
    keysByInsertOrder []string
}

func (sm *SortedMap) AddWord(word string) {
    if count, ok := sm.wordCount[word]; ok {
        sm.wordCount[word] = count + 1
    } else {
        sm.keysByInsertOrder = append(sm.keysByInsertOrder, word)
        sm.wordCount[word] = 1
    }
}

func (sm *SortedMap) GetWords(n, m int) {
    mostRepeated := make([]string, n, n)
    repeatedCounts := make([]int, n, n)
    for _, word := range sm.keysByInsertOrder {
        if len(word) <= m {
            continue
        }
        currCount := sm.wordCount[word]
        for i := 0; i < n; i++ {
            if currCount > repeatedCounts[i] {
                
            }
        }
    }
}