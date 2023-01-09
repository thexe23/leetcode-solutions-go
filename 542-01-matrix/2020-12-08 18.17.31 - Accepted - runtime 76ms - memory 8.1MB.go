func updateMatrix(matrix [][]int) [][]int {
    queue := [][]int{}
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                queue = append(queue, []int{i, j})
            }else {
                matrix[i][j] = 1<< 31 - 1
            }
        }
    }
    direct := [][]int{{0,1},{1,0}, {-1,0}, {0,-1}}
    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]
        for _, d := range direct {
            r := p[0] + d[0]
            c := p[1] + d[1]
            if r >= m || r < 0 || c >= n || c < 0 || matrix[r][c] < matrix[p[0]][p[1]] + 1 {
                continue
            }else {
                 matrix[r][c] = matrix[p[0]][p[1]] + 1
                queue = append(queue, []int{r,c})
            }
        }
    }
    return matrix
}