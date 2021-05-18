func minimumFromFour() int {
    var min int
    fmt.Scan(&min)
    for i := 0; i < 4 - 1; i++ {
        var tmp int
        fmt.Scan(&tmp)
        if tmp < min {
            min = tmp
        }
    }
    return min
}