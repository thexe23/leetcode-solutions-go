func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    temp := x
    p := x % 10
    reverse := 0
    for x > 0 {
        reverse = reverse * 10 + p
        x /= 10
        p = x % 10
    }
    if reverse == temp {
        return true
    }else{
        return false
    }
}