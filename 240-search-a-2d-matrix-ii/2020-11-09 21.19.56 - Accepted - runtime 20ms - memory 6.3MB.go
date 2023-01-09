func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    row, col := 0, len(matrix[0]) - 1
    for col >= 0 && row <= len(matrix) - 1{
        if matrix[row][col] == target {
            return true
        }else if target > matrix[row][col] {
            row++
        }else {
            col--
        }
    }
    return false
}