func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    start, end := 0, 0
    res := 0
    for start < len(s) && end < len(s) {
        start = max(start, m[s[end]])
        res = max(res, end - start + 1)
        m[s[end]] = end + 1
        end++
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}