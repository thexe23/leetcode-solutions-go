func plusOne(digits []int) []int {
    n := len(digits)
    digits[n - 1]++
    re := 0
    for i := n - 1; ; i--{
        digits[i] += re
        re = digits[i] / 10
        digits[i] %= 10
        if re == 0 || i <= 0 {
            break
        }
    }
    
    if re != 0 {
        res := []int{re}
        res = append(res, digits...)
        return res
    }
    return digits
}