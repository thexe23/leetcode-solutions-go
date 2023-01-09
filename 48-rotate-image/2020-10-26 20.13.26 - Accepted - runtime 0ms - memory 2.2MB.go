func rotate(matrix [][]int)  {
    n := len(matrix)
    
    for i := 0; i < (n + 1)/ 2; i++ {
        for j := 0; j < n /2; j++ {
           matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
        }
    }
}