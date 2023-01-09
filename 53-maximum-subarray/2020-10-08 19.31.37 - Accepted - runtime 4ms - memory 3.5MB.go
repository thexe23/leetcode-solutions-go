func maxSubArray(nums []int) int {
    N := len(nums)
    dp := make([]int, N, N)
    dp[0] = nums[0]
    res := dp[0]
    for i := 1; i < N; i++ {
        if dp[i - 1] > 0 {
            dp[i] = dp[i - 1] + nums[i]
        }else {
            dp[i] = nums[i]
        }
        res = max(res, dp[i])
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}