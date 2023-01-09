func minimumTotal(triangle [][]int) int {
    h := len(triangle)
    dp := make([]int, h)
    copy(dp, triangle[h - 1])
    for i := h - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j]
        }
    }
    return dp[0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}