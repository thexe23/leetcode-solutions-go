func minReorder(n int, connections [][]int) int {
    aj := make([][]int, n)
    
    for _, conn := range connections {
        aj[conn[0]] = append(aj[conn[0]], conn[1])
        aj[conn[1]] = append(aj[conn[1]], -conn[0])
    }
    visited := make([]bool, n)
    return dfs(aj, visited, 0)
}

func dfs(aj [][]int, visited []bool, from int) int{
    visited[from] = true
    change := 0
    for _, to := range aj[from] {
        if !visited[abs(to)] {
            change += dfs(aj, visited, abs(to))
            if to > 0 {
                change++
            }
        }
    }
    return change
}

func abs (x int) int {
    if x > 0 { return x}
    return -x
}