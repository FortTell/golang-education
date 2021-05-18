func sumInt(ints ...int) (int, int) {
    sum := 0
    for _, elem := range ints {
        sum += elem
    }
    return len(ints), sum
}