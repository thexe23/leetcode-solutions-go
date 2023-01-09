func isPalindrome(x int) bool {
    var rev, pop, current int
    
    if (x < 0 || (x % 10 == 0 && x != 0)) {
        return false
    }
    for current = x;current != 0; current /= 10 {
        pop = current % 10
        rev = rev*10 + pop
    }
    return x == rev
}