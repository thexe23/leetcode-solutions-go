func spiralOrder(matrix [][]int) []int {
	L := len(matrix)
	if L == 0 {
		return []int{}
	}
	W := len(matrix[0])
	rowBegin, rowEnd, colBegin, colEnd := 0, L - 1, 0, W - 1
	res := []int{}
	i, j := 0, 0
	for len(res) < L * W {
		for j = colBegin; j <= colEnd && len(res) < L * W; j++ {
			res = append(res, matrix[i][j])
		}
		j--
		rowBegin++
		for i = rowBegin; i <= rowEnd && len(res) < L * W; i++ {
			res = append(res, matrix[i][j])
		}
		i--
		colEnd--
		for j = colEnd; j >= colBegin && len(res) < L * W; j-- {
			res = append(res, matrix[i][j])
		}
		j++
		rowEnd--
		for i = rowEnd; i >= rowBegin && len(res) < L * W; i--{
			res = append(res, matrix[i][j])
		}
		i++
		colBegin++
	}
	return res
}