func canPartition(nums []int) bool {
    // calculate sum and max
    sum, max := 0, 0
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    // if sum is odd, can't be partitioned into two equal subsets, false
    if sum & 1 == 1 {
        return false
    }
    target := sum / 2
    // if max > target, sum of others must less than target, false
    if max > target {
        return false
    }
    // init dp
    dp := make([]bool, target + 1)
    dp[0] = true
    
    for _, v := range nums {
        for j := target; j >= v; j-- {
            dp[j] = dp[j] || dp[j - v]
        }
    }
    
    return dp[target]
}