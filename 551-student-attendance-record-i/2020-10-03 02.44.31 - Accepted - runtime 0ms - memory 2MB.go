func checkRecord(s string) bool {
    absent := 0
    for i := 0; i < len(s); i++ {
        if s[i] == byte('A') {
            absent++
            if absent > 1 {
                return false
            }
        }
        if s[i] == byte('L') && i < len(s) - 2 && s[i + 1] == byte('L')&&s[i + 2] == byte('L'){
            return false
        }
    }
    return true
}