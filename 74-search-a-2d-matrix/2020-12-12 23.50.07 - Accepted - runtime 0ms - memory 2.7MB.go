func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    start, end := 0, len(matrix) * len(matrix[0]) - 1
    for start <= end {
        mid := start + (end - start) / 2
        midRow := mid / len(matrix[0])
        midCol := mid - len(matrix[0]) * midRow
        if matrix[midRow][midCol] == target {
            return true
        }else if matrix[midRow][midCol] < target {
            start = mid + 1
        }else {
            end = mid - 1
        }
    }
    return false
}