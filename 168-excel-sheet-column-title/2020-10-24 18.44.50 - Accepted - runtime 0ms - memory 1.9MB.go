func convertToTitle(n int) string {
    res := ""
    for n > 0 {
        n--
        res = string(n % 26 + 'A') + res
        n /= 26
    }
    return res
}