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
    
    for dividend >= divisor {
        temp, bit := divisor, 1
        for dividend > temp<< 1{
            temp = temp<< 1
            bit = bit<<1
        }
        dividend -= temp
        res += bit
    }
    if sign  == 1 {
        return res
    }
    return -res
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}