func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := 1; i < n; i++ {
        dp[i] = i
        for j := 0; j < i; j++ {
            if nums[j] + j >= i {
                dp[i] = min(dp[i], dp[j] + 1)
            }
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