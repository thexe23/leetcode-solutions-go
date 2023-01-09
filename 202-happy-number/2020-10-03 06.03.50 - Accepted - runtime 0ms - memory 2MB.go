func isHappy(n int) bool {
    slow, fast := n, n
    for{
        fast = getNum(getNum(fast))
        slow = getNum(slow)
        
         if fast == 1 {
            return true
        }
        
        if fast == slow {
            break
        }
    }
    return false
}

func getNum(n int) int {
    sum := 0
    for n != 0 {
        sum += (n % 10) * (n % 10)
        n = n / 10
    }
    return sum
}