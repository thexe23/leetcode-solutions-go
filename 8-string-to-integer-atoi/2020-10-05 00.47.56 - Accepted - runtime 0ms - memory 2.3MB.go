func myAtoi(s string) int {
    IntMax := 1<<31 - 1
    IntMin := -1<<31
    i := 0
    res := 0
    // Skip the whitespace
    for i < len(s) && s[i] == ' '{
        i++
    }
    
    // Check if there is a sign before number
    sign := 1
    if i < len(s) && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }
    
    // Get the number
    for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
        if res > IntMax / 10 || (res == IntMax / 10 && int(s[i] - '0') > IntMax % 10){
            if sign == 1 {
                return IntMax
            }
            return IntMin
        }
        res = res * 10 + int(s[i] - '0')
        i++
    }
    return sign * res
}