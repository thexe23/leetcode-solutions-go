func divide(dividend int, divisor int) int {
    res := 0
    sign := 1
    if dividend == -1<<31 && divisor == -1 {
        return 1<<31 - 1
    }
    if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        sign = -1
    }
    dividend = abs(dividend)
    divisor = abs(divisor)
    for dividend >= 0 {
        dividend -= divisor
        res++
    }
    if sign  == 1 {
        return res - 1
    }
    return -res + 1
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}