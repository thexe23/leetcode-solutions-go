func uniquePaths(m int, n int) int {
    dp := make([][]int, m + 1)
    for k := 0; k < m + 1; k++ {
        dp[k] = make([]int, n + 1)
    }
    dp[1][0] = 1
    dp[0][1] = 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1] 
        }
    }
    return dp[m][n]
}

// dp[m, n] = dp[m-1, n] + dp[m, n - 1]
// dp[1, 1] = 1
// dp[1, 2] = dp[1, 1] + dp[0, 2]
// dp[2, 1] = dp[1, 1] + dp[2, 0]
//