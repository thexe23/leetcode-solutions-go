func lengthOfLongestSubstring(s string) int {
    end, res:= 0, 0
    m := make(map[byte]int)
    
    for start := -1; start < len(s); start++ {
        if start != -1 {
            m[s[start]]--
        }
        for end < len(s) && m[s[end]] == 0 {
            m[s[end]]++
            res = max(res, end - start )
            end++
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }else {
        return y
    }
}