func generateParenthesis(n int) []string {
    res := []string{}
    backtrack(&res, n, n, []byte{})
    return res
}

func backtrack(res *[]string, left, right int, row []byte) {
    if left == 0 && right == 0 {
        *res = append(*res, string(row))
        return
    }
    
    if left > 0 {
        row = append(row, '(')
        backtrack(res, left - 1, right, row)
        row = row[:len(row) - 1]
    }
    
    if right > left {
        row = append(row, ')')
        backtrack(res, left, right - 1, row)
        row = row[:len(row) - 1]
    }
}