package main

import (
    "fmt"
    //"strconv"
    "btree/src/btree"
)

func main() {
    /*var a string
    fmt.Scan(&a)
    aa, _ := strconv.Atoi(a[1:])
    fmt.Printf("%T\n", a[:1])
    fmt.Printf("%t\n", a[:1] == "+")
    fmt.Printf("%d\n", aa)*/
    tree := btree.ParseAndMakeTree()
    fmt.Println(tree.Traverse())
}