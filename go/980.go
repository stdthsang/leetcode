func uniquePathsIII(grid [][]int) int {
    x, y, total := 0, 0, 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                x, y = i, j
            }
            if grid[i][j] != -1 {
                total++
            }
        }
    }
    return dfs(grid, x, y, 1, total)
}

func dfs(grid [][]int, x, y int, cnt, total int) int {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == -1 {
        return 0
    }
    if grid[x][y] == 2 {
        if cnt == total {
            return 1
        }
        return 0
    }
    grid[x][y] = -1
    pathCnt := dfs(grid, x+1, y, cnt+1, total) + dfs(grid, x, y+1, cnt+1, total) +
        dfs(grid, x-1, y, cnt+1, total) + dfs(grid, x, y-1, cnt+1, total)
    grid[x][y] = 0
    return pathCnt
}

