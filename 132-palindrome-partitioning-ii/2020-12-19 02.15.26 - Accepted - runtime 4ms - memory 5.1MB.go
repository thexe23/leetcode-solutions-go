func minCut(s string) int {
    if len(s) == 1 {
        return 0
    }
    n := len(s)
    pal := make([][]bool, n)
    for i := 0; i < n; i++ {
        pal[i] = make([]bool, n)
    }
    dp := make([]int, n)
    for i := 0; i < len(s); i++ {
        dp[i] = i
        for j := 0; j <= i; j++ {
            if s[j] == s[i] && (j + 1 >= i - 1 || pal[j + 1][i - 1]) {
                pal[j][i] = true
                if j == 0 {
                    dp[i] = 0
                }else {
                    dp[i] = min(dp[i], dp[j - 1] + 1)
                }
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