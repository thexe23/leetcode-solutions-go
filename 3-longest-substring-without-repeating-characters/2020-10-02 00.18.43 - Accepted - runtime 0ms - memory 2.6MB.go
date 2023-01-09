func lengthOfLongestSubstring(s string) int {
    res := 0
    smap := make([]int, 128)
    for i, j := 0, 0; j < len(s); j++{
        i = max(smap[s[j]], i)
        res = max(res, j - i + 1)
        smap[s[j]] = j + 1
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}