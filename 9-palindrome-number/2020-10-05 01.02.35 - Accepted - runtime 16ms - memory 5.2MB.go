func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    re := 0
    for x > re {
        re = re * 10 + x % 10
        x /= 10
    }
    if x == re || x == re / 10{
        return true
    }
    
    return false
}