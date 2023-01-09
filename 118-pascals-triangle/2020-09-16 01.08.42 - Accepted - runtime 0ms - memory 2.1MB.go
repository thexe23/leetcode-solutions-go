func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    
     if numRows == 0 {
        return res
    }
    
    res[0] = []int{1}
    
    for i := 1; i < numRows; i++ {
        res[i] = make([]int, 0)
        res[i] = append(res[i], 1)
        for j := 1; j < i; j++ {
            res[i] = append(res[i], res[i-1][j-1] + res[i-1][j])
        }
        res[i] = append(res[i], 1)
    }
    return res
}