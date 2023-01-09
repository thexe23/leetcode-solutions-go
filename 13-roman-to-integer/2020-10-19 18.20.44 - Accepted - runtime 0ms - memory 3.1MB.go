func romanToInt(s string) int {
    m := map[byte]int{
        'M': 1000,
        'D': 500,
        'C':100,
        'L':50,
        'X':10,
        'V':5,
        'I':1,
    }
    res := m[s[len(s) - 1]]
    for i := len(s) - 2; i >= 0; i-- {
        if m[s[i]] >= m[s[i + 1]] {
            res += m[s[i]]
        }else {
            res -= m[s[i]]
        }
    }
    return res
}