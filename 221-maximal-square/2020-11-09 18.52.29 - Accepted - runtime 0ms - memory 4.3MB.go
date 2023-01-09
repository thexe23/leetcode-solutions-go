func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    h, w := len(matrix), len(matrix[0])
    dp := make([][]int, h)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, w)
    }
    res := 0
    for i := 0; i < h; i++ {
        for j := 0; j < w; j++ {
            if i == 0 || j == 0 || matrix[i][j] == '0'{
                dp[i][j] = int(matrix[i][j] - '0')
            }else {
               dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]) + 1
            }
            if dp[i][j] > res {
                res = dp[i][j]
            }
        }
    }
    return res * res
}

func min(x, y, z int) int {
    if x < y {
        if x < z {
            return x
        }
        return z
    }
    if y < z {
        return y
    }
    return z
}