func minWindow(s string, t string) string {
    m := make([]int, 128)
    for _, v := range t {
        m[v]++
    }
    
    left, right, res := 0, 0, ""
    count := len(t)
    
    for right < len(s) {
        if m[s[right]] > 0 {
            count--
        }
        m[s[right]]--
        right++
        
        for count == 0 {
            if res == "" || right - left < len(res) {
                res = s[left : right]
            }
            m[s[left]]++
            if m[s[left]] > 0 {
                count++
            }
            left++
        }
    }
    return res
}