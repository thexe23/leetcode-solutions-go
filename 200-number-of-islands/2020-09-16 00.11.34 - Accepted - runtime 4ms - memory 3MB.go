func numIslands(grid [][]byte) int {
    var count int = 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                count += 1
                Travel(grid, i, j)
            }
        }
    }
    return count
}

func Travel(grid [][]byte, i int, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    Travel(grid, i+1, j) // Right
    Travel(grid, i, j+1) // Down
    Travel(grid, i-1, j) // Left
    Travel(grid, i, j-1) // Up
}