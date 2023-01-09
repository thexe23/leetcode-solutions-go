func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {
        return []int{}
    }
    need, have := [26]byte{}, [26]byte{}
    res := []int{}
    for i := 0; i < len(p); i++ {
        need[p[i] - 'a']++
        have[s[i] - 'a']++
    }
    var tail int
    for i := len(p); i < len(s); i++ {
        if need == have {
            res = append(res, tail)
        }
        have[s[i] - 'a']++
        have[s[tail] - 'a']--
        tail++
    }
    if have == need {
        res = append(res, tail)
    }
    return res
}