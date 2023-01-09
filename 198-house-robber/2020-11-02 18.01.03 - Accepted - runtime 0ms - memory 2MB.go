func rob(nums []int) int {
    dp := make([]int, len(nums) + 2)
    i := 2
    for i < len(nums) + 2 {
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i - 2])
        i++
    }
    return dp[i - 1]
}

//dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])

func max(x, y int) int {
    if x  > y {
        return x
    }
    return y
}