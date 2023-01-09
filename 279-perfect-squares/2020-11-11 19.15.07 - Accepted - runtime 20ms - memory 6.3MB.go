func numSquares(n int) int {
    dp := []int{0}
    
    for len(dp) <= n{
        l := len(dp)
        res := n
        for j := 1; j * j <= l; j++ {
            res = min(res, dp[l - j * j] + 1)
        }
        dp = append(dp, res)
    }
    return dp[n]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}