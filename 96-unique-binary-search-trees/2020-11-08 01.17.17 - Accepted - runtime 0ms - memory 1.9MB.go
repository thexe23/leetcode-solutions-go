func numTrees(n int) int {
    G := make([]int, n + 1)
    G[0] = 1
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            G[i] += G[j] * G[i - 1 - j]
        }
    }
    return G[n]
}