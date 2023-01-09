func totalNQueens(n int) int {
    count := 0
    var put func (n, row, cm, dm1, dm2 int)
    put = func (n, row, cm, dm1, dm2 int) {
    if row == n {
        count++
        return
    }
    
        for col := 1<< (n - 1); col > 0; col >>= 1 {
        if col & (cm|dm1|dm2) == 0 {
            put(n, row + 1, cm|col, (dm1|col)<<1, (dm2|col)>>1)
        }
    }
}
    put(n, 0, 0, 0, 0)
    return count
}