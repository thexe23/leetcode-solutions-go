func wordBreak(s string, wordDict []string) bool {
    wordMap := make(map[string]bool)
    dp := make([]bool, len(s) + 1)
    
    for _, v := range wordDict {
        wordMap[v] = true
    }
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := i - 1; j >= 0; j-- {
            if dp[j] && wordMap[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}