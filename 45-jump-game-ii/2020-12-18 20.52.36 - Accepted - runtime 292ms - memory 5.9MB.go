func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := 1; i < n; i++ {
        dp[i] = n
    }
    for i, v := range nums {
        for j := 1; j <= v && j < n - i; j++ {
            dp[i + j] = min(dp[i + j], dp[i] + 1)
        }
    }
    return dp[n - 1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}