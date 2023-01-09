func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    re := 0
    temp := x
    for temp != 0 {
        re = re * 10 + temp % 10
        temp /= 10
    }
    return re == x
}