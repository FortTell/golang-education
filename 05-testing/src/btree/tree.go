package btree

import (
    "fmt"
    "strconv"
)

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

type Node struct {
    Left *Node
    Right *Node
    Parent *Node
    Value int
    ValueCount int
}

func (node *Node) Height() int {
    if node == nil {
        return 0
    } else {
        return max(node.Left.Height(), node.Right.Height()) + 1
    }
} 

func (node *Node) BalanceFactor() int {
    if node == nil {
        return 0
    } else {
        return node.Left.Height() - node.Right.Height()
    }
}

func (node *Node) minValueNode() *Node {
    curr := node
    for curr.Left != nil {
        curr = curr.Left
    }
    return curr
}

func (node *Node) rotateLeft() *Node {
    pivot := node.Right
    tmp := pivot.Left
    
    pivot.Left = node
    pivot.Parent = node.Parent
    node.Parent = pivot
    
    node.Right = tmp
    if tmp != nil {
        tmp.Parent = node
    }
    
    if pivot.Parent != nil {
        if pivot.Parent.Left == node {
            pivot.Parent.Left = pivot
        } else {
            pivot.Parent.Right = pivot
        }
    }
    
    return pivot
}

func (node *Node) rotateRight() *Node {
    pivot := node.Left
    tmp := pivot.Right
    
    pivot.Right = node
    pivot.Parent = node.Parent
    node.Parent = pivot
    
    node.Left = tmp
    if tmp != nil {
        tmp.Parent = node
    }
    
    if pivot.Parent != nil {
        if pivot.Parent.Left == node {
            pivot.Parent.Left = pivot
        } else {
            pivot.Parent.Right = pivot
        }
    }
    
    return pivot
}

func (node *Node) rebalance() *Node {    
    bf := node.BalanceFactor()
    switch bf {
        case 2: // left deeper than right
            if node.Left.BalanceFactor() < 0 { // L-R
                node.Left = node.Left.rotateLeft()
            }
            return node.rotateRight() // L-L
        case -2: // right deeper than left
            if node.Right.BalanceFactor() > 0 { // R-L
                node.Right = node.Right.rotateRight()
            }
            return node.rotateLeft() //R-R
        default:
            return node
    }
}

func (node *Node) Insert(value int) *Node {
    if node == nil {
        newNode := Node{Value: value, ValueCount: 1}
        return &newNode
    }
    
    switch {
        case value < node.Value:
            leftSubRoot := node.Left.Insert(value)
            node.Left = leftSubRoot
            leftSubRoot.Parent = node
        case value > node.Value:
            rightSubRoot := node.Right.Insert(value)
            node.Right = rightSubRoot
            rightSubRoot.Parent = node
        default:
            node.ValueCount++
            return node
    }
    
    return node.rebalance()
}

func (node *Node) deleteNodeWithOneOrZeroChildren() *Node {
    var temp *Node
    switch {
        case node.Left == nil && node.Right == nil:
            return temp
        case node.Left != nil:
            return node.Left
        case node.Right != nil:
            return node.Right
        default:
            panic("Should not be here")
    }
}

func (node *Node) Delete(value int) *Node {
    if node == nil {
        return node
    }
    
    switch {
        case value < node.Value:
            node.Left = node.Left.Delete(value)
        case value > node.Value:
            node.Right = node.Right.Delete(value)
        default:
            node.ValueCount--
            if node.ValueCount > 0 {
                return node
            }
        
            if node.Left != nil && node.Right != nil {
                temp := node.Right.minValueNode()
                node.Value = temp.Value
                node.Right = node.Right.Delete(temp.Value)
            } else {
                node = node.deleteNodeWithOneOrZeroChildren()
            }
    }
    
    return node.rebalance()
}

func (node *Node) traverse(visitedNodes []*Node) []*Node {
    if node == nil {
        return visitedNodes
    }
    
    afterLeftSubtree := node.Left.traverse(visitedNodes)
    afterNode := append(afterLeftSubtree, node)
    afterRightSubtree := node.Right.traverse(afterNode)
    return afterRightSubtree
}

type AvlTree struct {
    Root *Node
}

func (tree *AvlTree) Insert(value int) {
    tree.Root = tree.Root.Insert(value)
}

func (tree *AvlTree) Delete(value int) {
    tree.Root = tree.Root.Delete(value)
}

func (tree *AvlTree) Traverse() []int {
    if tree.Root == nil {
        return []int{}
    }

    visitedNodes := make([]*Node, 0)
    orderedNodes := tree.Root.traverse(visitedNodes)
    var nodeValues []int
    for _, node := range orderedNodes {
        for i := 0; i < node.ValueCount; i++ {
            nodeValues = append(nodeValues, node.Value)
        }
    }
    return nodeValues
}

func ParseAndMakeTree() *AvlTree {
    var curr string
    tree := new(AvlTree)
    for fmt.Scan(&curr); true; fmt.Scan(&curr) {
        switch curr[:1] {
            case "+": 
                value, _ := strconv.Atoi(curr[1:])
                tree.Insert(value)
            case "-": 
                value, _ := strconv.Atoi(curr[1:])
                tree.Delete(value)
            default: return tree
        }
    }
    return tree
}