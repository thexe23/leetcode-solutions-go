func exist(board [][]byte, word string) bool {
    h, w := len(board), len(board[0])
    
    var check func(i, j, l int) bool
    check = func(i, j, l int) bool {
        if l == len(word) {
            return true
        }
        if i < 0 || i == h || j < 0 || j == w || board[i][j] != word[l] {
            return false
        }
        board[i][j] = '*'
        found := check(i + 1, j, l + 1) || check(i, j + 1, l + 1) || check(i - 1, j, l + 1) || check(i, j - 1, l + 1)
        board[i][j] = word[l]
        return found
    }
    
    for i := 0; i < h; i++ {
        for j := 0; j < w; j++ {
            if check(i, j, 0) {
                return true
            }
        }
    }
    return false
}