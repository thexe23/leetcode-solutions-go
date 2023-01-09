func generateParenthesis(n int) []string {
    res := []string{}
    findAnswer(&res, "", 0, 0, n)
    return res
}

func findAnswer(res *[]string, cur string, left, right, max int) {
    if len(cur) == max * 2 {
        *res = append(*res, cur)
        return
    }
    
    if left < max {
        findAnswer(res, cur + "(", left + 1, right, max)
    }
    
    if right < left{
        findAnswer(res, cur + ")", left, right + 1, max)
    }
}