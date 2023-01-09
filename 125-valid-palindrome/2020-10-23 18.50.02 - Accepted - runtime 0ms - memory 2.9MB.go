func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    start, end := 0, len(s) - 1
    for start < end {
        if !isAlphanumeric(s[start]) {
            start++
        }else if !isAlphanumeric(s[end]) {
            end--
        }else if s[start] == s[end] {
            start++
            end--
        }else {
            return false
        }
    }
    return true
}

func isAlphanumeric(s byte) bool {
    if s >= 'a' && s <= 'z' {
        return true
    }else if s >= 'A' && s <= 'Z' {
        return true
    }else if s >= '0' && s <= '9' {
        return true
    }
    return false
}