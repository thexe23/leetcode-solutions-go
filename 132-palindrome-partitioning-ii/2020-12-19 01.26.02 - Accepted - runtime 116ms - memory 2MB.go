func minCut(s string) int {
    if len(s) < 2 {
        return 0
    }
    n := len(s)
    dp := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = i - 1
        for j := 0; j < i; j++ {
            if isPalindrome(s, j, i - 1) {
                dp[i] = min(dp[i], dp[j] + 1)
            }
        }
    }
    return dp[n]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func isPalindrome(s string, i, j int) bool {
    for i < j {
        if s[i] == s[j] {
            i++
            j--
        }else {
            return false
        }
    }
    return true
}