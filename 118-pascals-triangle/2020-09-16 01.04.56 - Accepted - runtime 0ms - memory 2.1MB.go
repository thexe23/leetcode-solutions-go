func generate(numRows int) [][]int {
    res := make([][]int, numRows)
     if numRows == 0 {
        return res
    }
    res[0] = make([]int, 1)
    res[0][0] = 1
    
    for i := 1; i < numRows; i++ {
        res[i] = make([]int, 0)
        res[i] = append(res[i], 1)
        var prev_row = res[i-1]
        for j := 1; j < i; j++ {
            res[i] = append(res[i], prev_row[j-1] + prev_row[j])
        }
        res[i] = append(res[i], 1)
    }
    return res
}