func longestValidParentheses(s string) int {
    left, right, res := 0, 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        }else {
            right++
        }
        
        if left == right {
            res = max(res, right * 2)
        }else if right > left {
            right, left = 0, 0
        }
    }
    
    left, right = 0, 0
    
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        }else {
            right++
        }
        
        if left == right {
            res = max(res, right * 2)
        }else if right < left {
            right, left = 0, 0
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}