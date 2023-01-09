func reverse(x int) int {
    IntMax := 1 << 31 - 1
    IntMin := -1 << 31
    
    res := 0
    for x != 0 {
        old := x % 10
        if res > IntMax/10 || res < IntMin / 10 {
            return 0
        }
        res = res * 10 + old
        x = x / 10
    }
    return res
}