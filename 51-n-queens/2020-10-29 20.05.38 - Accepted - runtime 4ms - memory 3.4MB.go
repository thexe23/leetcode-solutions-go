func solveNQueens(n int) [][]string {
    res := [][]string{}
    chess := make([][]byte, n)
    for i := 0; i < n; i++ {
        chess[i] = make([]byte, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            chess[i][j] = '.'
        }
    }
    backtrack(chess, &res, 0)
    return res
}

func backtrack(chess [][]byte, res *[][]string, row int) {
    if row == len(chess) {
        *res = append(*res, toString(chess))
        return
    }
    
    for i := 0; i < len(chess); i++ {
        if check(chess, row, i) {
            chess[row][i] = 'Q'
            backtrack(chess, res, row + 1)
            chess[row][i] = '.'
        }
    }
}

func check(chess [][]byte, row, col int) bool {
    // check col
    for i := 0; i < row; i++ {
        if chess[i][col] == 'Q' {
            return false
        }
    }
    
    // check 45 degree
    for i, j := row, col; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
        if chess[i][j] == 'Q' {
            return false
        }
    }
    
    // check 135 degree
    for i, j := row, col; i >= 0 && j < len(chess); i, j = i - 1, j + 1 {
        if chess[i][j] == 'Q' {
            return false
        }
    }
    return true
}

func toString(chess [][]byte) []string {
    res := []string{}
    for i := 0; i < len(chess); i++ {
        str := []byte{}
        for j := 0; j < len(chess); j++ {
            str = append(str, chess[i][j]) 
        }
        res = append(res, string(str))
    }
    return res
}