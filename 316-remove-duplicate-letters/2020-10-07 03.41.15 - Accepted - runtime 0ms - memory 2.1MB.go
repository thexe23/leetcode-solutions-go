func removeDuplicateLetters(s string) string {
    lastInstance := [26]int{}
    seen := [26]bool{}
    res := []rune{}
    
    for i, r := range s {
        lastInstance[r - 'a'] = i
    }
    
    for i, r := range s {
        for len(res) > 0 {
            last := res[len(res) - 1]
            
            if r <= last && lastInstance[last - 'a'] >= i && !seen[r - 'a'] {
                res = res[:len(res) - 1]
                seen[last - 'a'] = false
                continue
            }
            break
        }
        
        if !seen[r - 'a'] {
            res = append(res, r)
            seen[r - 'a'] = true
        }
    }
    return string(res)
}