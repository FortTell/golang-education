func fibonacci(n int) int {
    switch n {
        case 0: return 0
        case 1, 2: return 1
        default: return fibonacci(n - 1) + fibonacci(n - 2)
    }
}