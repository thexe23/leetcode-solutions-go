func longestPalindrome(s string) string {
    if len(s) < 1 {
        return s
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := expand(s, i, i)
        len2 := expand(s, i, i+1)
        max_len := max(len1, len2)
        if max_len > (end - start) {
            start = i - (max_len - 1)/2
            end = i + max_len/2
        }
    }
    return s[start:end+1]
}

func expand(s string, left int, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(x,y int) int {
    if x > y {
        return x
    }else{
        return y
    }
}