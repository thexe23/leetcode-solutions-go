func isPalindrome(x int) bool {
    if x < 0 || ( x % 10 == 0 && x != 0 ){
        return false
    }
    p := x % 10
    reverse := 0
    for x > reverse {
        reverse = reverse * 10 + p
        x /= 10
        p = x % 10
    }
    if reverse == x || x == reverse/10 {
        return true
    }else{
        return false
    }
}