package main

import (
    "fmt"
    "btree/src/btree"
)

func main() {
    tree := btree.ParseAndMakeTree()
    fmt.Println(tree.Traverse())
}