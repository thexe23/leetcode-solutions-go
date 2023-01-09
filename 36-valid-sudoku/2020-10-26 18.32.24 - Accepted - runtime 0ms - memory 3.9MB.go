func isValidSudoku(board [][]byte) bool {
    rowMap := make([][]int, 10)
    colMap := make([][]int, 10)
    boxMap := make([][]int, 10)
    
    for i := 0; i < 9; i++ {
        rowMap[i] = make([]int, 10)
        colMap[i] = make([]int, 10)
        boxMap[i] = make([]int, 10)
    }
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] != '.'{
                if rowMap[i][board[i][j] - '0'] == 1 {
                    return false
                }
                if colMap[j][board[i][j] - '0'] == 1 {
                    return false
                }
                if boxMap[(i / 3) * 3 + j / 3][board[i][j] -'0'] == 1 {
                    return false
                }
                rowMap[i][board[i][j] - '0']++
                colMap[j][board[i][j] - '0']++
                boxMap[(i / 3) * 3 + j / 3][board[i][j] -'0']++
            }
        }
    }
    return true
}