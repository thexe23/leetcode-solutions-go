func countSubstrings(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        res += expand(s, i, i)
        res += expand(s, i, i + 1)
    }
    return res
}

func expand(s string, i, j int) int {
    res := 0
    for i >= 0 && j < len(s) && s[i] == s[j] {
        res++
        i--
        j++
    }
    return res
}