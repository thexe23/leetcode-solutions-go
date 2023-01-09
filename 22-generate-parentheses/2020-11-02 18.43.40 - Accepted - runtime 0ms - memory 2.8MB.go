func generateParenthesis(n int) []string {
    res := []string{}
    backtrack(&res, n, n, "")
    return res
}

func backtrack(res *[]string, left, right int, row string) {
    if left == 0 && right == 0 {
        *res = append(*res, row)
        return
    }
    
    if left > 0 {
        backtrack(res, left - 1, right, row + "(")
    }
    
    if right > left {
        backtrack(res, left, right - 1, row + ")")
    }
}