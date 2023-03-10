func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for k := 0; k < m; k++ {
        dp[k] = make([]int, n)
        dp[k][0] = 1
    }
    for k := 0; k < n; k++ {
        dp[0][k] = 1
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1] 
        }
    }
    return dp[m -1 ][n - 1]
}