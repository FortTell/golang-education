package sortedSlice

type SortedSlice struct {
    slice []int
}

func (tree *SortedSlice) Insert(value int) {
    isInserted := false
    newSlice := make([]int, len(tree.slice) + 1)
    for i := 0; i < len(tree.slice); i++ {
        var insertIndex int
        if isInserted { 
            insertIndex = i + 1
        } else {
            insertIndex = i
        }
        
        newSlice[insertIndex] = tree.slice[i]
        if !isInserted && value > tree.slice[i] {
            isInserted = true
            newSlice[i + 1] = value
        }
    }
    tree.slice = newSlice
}

func (tree *SortedSlice) Delete(value int) {
    deleteIndex := -1
    for i := 0; i < len(tree.slice); i++ {
        if tree.slice[i] == value {
            deleteIndex = i
            break
        }
    }
    if deleteIndex == -1 {
        return
    }
    
    tree.slice = append(tree.slice[:deleteIndex], tree.slice[deleteIndex + 1:]...)
}

func (tree *SortedSlice) Traverse() []int {
    values := make([]int, len(tree.slice))
    for i := 0; i < len(tree.slice); i++ {
        values[i] = tree.slice[i]
    }
    return values
}

func ParseAndMakeSortedSlice() *SortedSlice {
    var curr string
    sortedSlice := new(SortedSlice)
    for fmt.Scan(&curr); true; fmt.Scan(&curr) {
        switch curr[:1] {
            case "+": 
                value, _ := strconv.Atoi(curr[1:])
                sortedSlice.Insert(value)
            case "-": 
                value, _ := strconv.Atoi(curr[1:])
                sortedSlice.Delete(value)
            default: return sortedSlice
        }
    }
    return sortedSlice
}