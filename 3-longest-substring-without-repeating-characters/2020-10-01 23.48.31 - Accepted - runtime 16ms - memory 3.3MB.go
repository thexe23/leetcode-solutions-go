func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    start := 0
    end := 0
    res := 0
    for ;start < len(s); start++ {
        if start != 0 {
             m[s[start - 1]]--
        }
        for ;end < len(s) && m[s[end]] == 0; end++ {
            m[s[end]]++
            res = max(res, end - start + 1)
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}