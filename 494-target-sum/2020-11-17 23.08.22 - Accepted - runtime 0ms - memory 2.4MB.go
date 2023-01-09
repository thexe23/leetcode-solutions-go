func findTargetSumWays(nums []int, S int) int {
// positiveSubSet + negativeSubSet = S
// sumP + sumN = S
// sumP + sumN + sumP - sumN = S + sum
// 2 * sumP = sum + S
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum < S || (sum + S) & 1 == 1 {
        return 0
    }
    target := (sum + S) / 2
    dp := make([]int, target + 1)
    dp[0] = 1
    for _, v := range nums {
        for j := target; j >= v; j-- {
            dp[j] = dp[j] + dp[j - v]
        }
    }
    return dp[target]
}