func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := 1; i < amount + 1; i++ {
        dp[i] = amount + 1
        for _, v := range coins {
            if v <= i {
                 dp[i] = min(dp[i], dp[i - v] + 1)
            }
        }
    }
    
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}