func partitionLabels(S string) []int {
    last := [26]int{}
    res := []int{}
    for i := 0; i < len(S); i++ {
        last[S[i] - 'a'] = i
    }
    bound := 0
    anchor := 0
    for i := 0; i < len(S); i++ {
        bound = max(bound, last[S[i] - 'a'])
        if i == bound {
            res = append(res, i - anchor + 1)
            anchor = i + 1
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