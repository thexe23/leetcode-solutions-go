func longestPalindrome(s string) string {
    Max := 0
    res := ""
    for i := 0 ; i < len(s); i++ {
        s1 := expand(s, i, i)
        s2 := expand(s, i, i + 1)
        if len(s1) > Max {
            Max = len(s1)
            res = s1
        }
        if len(s2) > Max {
            Max = len(s2)
            res = s2
        }
    }
    return res
}

func expand(s string, left, right int) string {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return s[left + 1: right]
}