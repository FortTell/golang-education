contains := func(s []rune, e rune) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

fn := func(src uint) uint {
    badDigits := []rune{'0','1','3','5','7','9'}
    digits := strconv.FormatUint(uint64(src), 10)
    goodDigits := make([]rune, 0, len(digits))
    for _, digit := range digits {
        if !contains(badDigits, digit) {
            goodDigits = append(goodDigits, digit)
        }
    }
    if len(goodDigits) == 0 {
        return uint(100)
    } else {
        result, _ := strconv.ParseUint(string(goodDigits), 10, 32)
        return uint(result)
    }
}