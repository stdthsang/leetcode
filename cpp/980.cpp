class Solution {
public:
    int uniquePathsIII(vector<vector<int>>& grid) {
        int x = 0, y = 0, total = 0;
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[0].size(); ++j) {
                if (grid[i][j] == 1) {
                    x = i;
                    y = j;
                }
                if (grid[i][j] != -1)
                    ++total;
            }
        }
        return dfs(grid, x, y, 1, total);
    }
    int dfs(vector<vector<int>>& grid, int x, int y, int cnt, int total) {
        if (x < 0 || y < 0 || x >= grid.size() || y >= grid[0].size() || grid[x][y] == -1)
            return 0;
        if (grid[x][y] == 2)
            return cnt == total;
        grid[x][y] = -1;
        int pathCnt = dfs(grid, x + 1, y, cnt + 1, total) + dfs(grid, x - 1, y, cnt + 1, total) +
                dfs(grid, x, y + 1, cnt + 1, total) + dfs(grid, x, y - 1, cnt + 1, total);
        grid[x][y] = 0;
        return pathCnt;
    }
};
