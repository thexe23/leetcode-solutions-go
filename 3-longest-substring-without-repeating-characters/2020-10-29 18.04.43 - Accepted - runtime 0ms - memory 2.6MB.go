func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    charMap := make([]int, 128)
    res := 0
    for i, j := 0, 0; j < len(s); j++ {
        i = max(charMap[s[j]], i)
        res = max(res, j - i + 1)
        charMap[s[j]] = j + 1
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}