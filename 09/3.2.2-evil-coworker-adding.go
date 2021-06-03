func adding(s1 string, s2 string) int64 {
    s1Cleaned := keepDigits(s1)
    s2Cleaned := keepDigits(s2)
    a, _ := strconv.Atoi(s1Cleaned)
    b, _ := strconv.Atoi(s2Cleaned)
    return int64(a + b)
}

func keepDigits(s string) string {
    cleaned := make([]rune, 0, len(s))
    for _, r := range s {
        if unicode.IsDigit(r) {
            cleaned = append(cleaned, r)
        }
    }
    return string(cleaned)
}
