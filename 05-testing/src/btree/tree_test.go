package btree

import (
    "testing"
)

func ShouldEqual(t *testing.T, actual interface{}, expected interface{}) {
    if actual != expected {
        t.Errorf("expected %#v, but was %#v", expected, actual)
    }
}

func TestNilNodeHeight(t *testing.T) {
    var root *Node
    got := root.Height()
    if got != 0 {
        t.Errorf("nil node should have height 0 but was %d", got)
    }
}

func TestNodeHeight(t *testing.T) {
    root := Node{Value: 1}
    got := (&root).Height()
    if got != 1 {
        t.Errorf("leaf node should have height 1 but was %d", got)
    }
    otherNode := Node{Value: 2}
    root.Right = &otherNode
    got = (&root).Height()
    if got != 2 {
        t.Errorf("should have height 2 but was %d", got)
    }
    leftNode1 := Node{Value: -10}
    root.Left = &leftNode1
    leftNode2 := Node{Value: -100}
    leftNode1.Left = &leftNode2
    got = (&root).Height()
    if got != 3 {
        t.Errorf("should have height 3 but was %d", got)
    }
}

func TestNodeInsertion(t *testing.T) {
    root := Node{Value: 1}
    (&root).Insert(2)
    if root.Right == nil || root.Right.Value != 2 {
        t.Errorf("should have inserted bigger value to the right")
    }
    (&root).Insert(0)
    if root.Left == nil || root.Left.Value != 0 {
        t.Errorf("should have inserted smaller value to the left")
    }
}

func TestLRNodeRebalancing(t *testing.T) {
    var nilNode *Node
    tree := new(AvlTree)
    tree.Insert(5)
    ShouldEqual(t, tree.Root.Value, 5)
    tree.Insert(3)
    ShouldEqual(t, tree.Root.Right, nilNode)
    ShouldEqual(t, tree.Root.Left.Value, 3)
    tree.Insert(4)
    ShouldEqual(t, tree.Root.Value, 4)
    ShouldEqual(t, tree.Root.Left.Value, 3)
    ShouldEqual(t, tree.Root.Right.Value, 5)
}

func TestDeleting(t *testing.T) {
    var nilNode *Node
    tree := new(AvlTree)
    for i := 0; i <= 6; i++ {
        tree.Insert(i)
    }
    ShouldEqual(t, tree.Root.Value, 3)
    ShouldEqual(t, tree.Root.Height(), 3)
    
    tree.Delete(6)
    ShouldEqual(t, tree.Root.Right.Value, 5)
    ShouldEqual(t, tree.Root.Right.Right, nilNode)
    
    tree.Delete(5)
    ShouldEqual(t, tree.Root.Right.Value, 4)
    ShouldEqual(t, tree.Root.Right.Left, nilNode)
    ShouldEqual(t, tree.Root.Right.Right, nilNode)
    
    tree.Delete(4)
    ShouldEqual(t, tree.Root.Value, 1)
    ShouldEqual(t, tree.Root.Left.Value, 0)
    ShouldEqual(t, tree.Root.Right.Value, 3)
    ShouldEqual(t, tree.Root.Right.Left.Value, 2)
}

func TestTraversing(t *testing.T) {
    tree := new(AvlTree)
    expectedValues := make([]int, 7)
    for i := 0; i <= 6; i++ {
        expectedValues[i] = i
        tree.Insert(i)
    }
    values := tree.Traverse()
    for i := 0; i <= 6; i++ {
        ShouldEqual(t, values[i], expectedValues[i])
    }
}

func TestRepeatedValueInsert(t *testing.T) {
    tree := new(AvlTree)
    tree.Insert(4)
    tree.Insert(4)
    
    var nilNode *Node
    ShouldEqual(t, tree.Root.Value, 4)
    ShouldEqual(t, tree.Root.ValueCount, 2)
    ShouldEqual(t, tree.Root.Left, nilNode)
    ShouldEqual(t, tree.Root.Right, nilNode)
}

func TestRepeatedValueDelete(t *testing.T) {
    tree := new(AvlTree)
    tree.Insert(4)
    tree.Insert(4)
    tree.Delete(4)
    
    var nilNode *Node
    ShouldEqual(t, tree.Root.Value, 4)
    ShouldEqual(t, tree.Root.ValueCount, 1)
    ShouldEqual(t, tree.Root.Left, nilNode)
    ShouldEqual(t, tree.Root.Right, nilNode)
}

func TestRepeatedValueTotalDelete(t *testing.T) {
    tree := new(AvlTree)
    tree.Insert(4)
    tree.Insert(4)
    tree.Delete(4)
    tree.Delete(4)
    
    var nilNode *Node
    ShouldEqual(t, tree.Root, nilNode)
}

func TestRepeatedValuePrinting(t *testing.T) {
    valuesToInsert := [...]int{4, 4, 5, 6, 0}
    tree := new(AvlTree)
    for _, v := range valuesToInsert {
        tree.Insert(v)
    }
    
    expectedValues := [...]int{0, 4, 4, 5, 6}
    values := tree.Traverse()
    for i := 0; i < len(expectedValues); i++ {
        ShouldEqual(t, values[i], expectedValues[i])
    }
}