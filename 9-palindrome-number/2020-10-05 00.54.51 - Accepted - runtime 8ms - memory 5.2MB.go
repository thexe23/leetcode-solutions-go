func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    temp := x
    re := 0
    for x != 0 {
        re = re * 10 + x % 10
        x /= 10
    }
    if temp == re {
        return true
    }
    
    return false
}