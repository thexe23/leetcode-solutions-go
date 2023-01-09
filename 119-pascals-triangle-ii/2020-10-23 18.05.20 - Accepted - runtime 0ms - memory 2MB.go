func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    res := make([]int, rowIndex + 1)
    res[0] = 1
    res[rowIndex] = 1
    prev := getRow(rowIndex - 1)
    for i := 1; i < rowIndex; i++ {
        res[i] = prev[i] + prev[i - 1]
    }
    return res
}