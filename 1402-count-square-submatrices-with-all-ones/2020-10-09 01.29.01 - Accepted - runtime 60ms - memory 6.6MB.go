func countSquares(matrix [][]int) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m + 1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n + 1)
    }
    res := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if matrix[i - 1][j - 1] == 1 {
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1]) + 1
            }else {
                dp[i][j] = 0
            }
            res += dp[i][j]
        }
    }
    return res
}

func min(x, y, z int) int {
    min := x
    if y < min {
        min = y
    }
    
    if z < min {
        min = z
    }
    return min
}