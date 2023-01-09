func checkInclusion(s1 string, s2 string) bool {
    m := make([]int, 26)
    for _, v := range s1 {
        m[v - 'a']++
    }
    left, right, count := 0, 0, len(s1)
    for right < len(s2) {
        if m[s2[right] - 'a'] > 0 {
            count--
        }
        m[s2[right] - 'a']--
        right++
        
        if right >= len(s1) {
            if count == 0 {
                return true
            }
            m[s2[left] - 'a']++
            if m[s2[left] - 'a'] > 0 {
                count++
            }
            left++
        }
    }
    return false
}