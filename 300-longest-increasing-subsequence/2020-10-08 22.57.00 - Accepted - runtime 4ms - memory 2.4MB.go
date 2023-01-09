func lengthOfLIS(nums []int) int {
    N := len(nums)
    if N == 0 {
        return 0
    }
    dp := make([]int, N)
    dp[0] = 1
    res := 1
    for i := 1; i < N; i++ {
        m := 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                m = max(dp[j], m)
            }
        }
        dp[i] = m + 1
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