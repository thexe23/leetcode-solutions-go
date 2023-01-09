func minPathSum(grid [][]int) int {
    H := len(grid)
    W := len(grid[0])
    dp := make([][]int, H)
    for i := 0; i < H; i++ {
        dp[i] = make([]int, W)
    }
    dp[0][0] = grid[0][0]
    for i := 1; i < H; i++{
        dp[i][0] = dp[i - 1][0] + grid[i][0]
    }
    
    for i := 1; i < W; i++ {
        dp[0][i] = dp[0][i - 1] + grid[0][i]
    }
    
    
    for i := 1; i < H; i++ {
        for j := 1; j < W; j++ {
            dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
        }
    }
    return dp[H - 1][W - 1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
/*
    dp[i][j] = min(d[i - 1][j], d[i][j - 1]) + grid[i][j]
*/