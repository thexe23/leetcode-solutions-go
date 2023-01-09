func rotatedDigits(N int) int {
    count := 0
    for i := 1; i <= N; i++{
        num := i
        good, invalid, remain := 0, 0, 0
        for num > 0 {
            n := num % 10
            if n == 0 || n == 1 || n == 8{
                remain++
            }else if n == 2 || n == 5 || n == 6 || n == 9 {
                good++
            }else {
                invalid++
                break
            }
            num /= 10
        }
        if invalid == 0 && good != 0 {
            count++
        }
    }
    return count
}