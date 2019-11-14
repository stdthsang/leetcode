class Solution {
public:
    vector<vector<string>> solveNQueens(int n) {
        vector<string> cur(n, string(n,'.'));
        vector<vector<string>> res;
        helper(0, n, cur, res);
        return res;
    }
    
    void helper(int row, int n, vector<string> &cur, vector<vector<string>> &out) {
        if (row == n) {
            out.push_back(cur);
            return;
        }
        for (int i = 0; i < n; ++i) {
            if (isValid(row, i, n, cur)) {
                cur[row][i] = 'Q';
                helper(row + 1, n, cur, out);
                cur[row][i] = '.';
            }
        }
    }
    
    bool isValid(int row, int col, int n, vector<string> &cur) {
        for (int i = 0; i < row; ++i)
            if (cur[i][col] == 'Q')
                return false;
        for (int i = 0; i < col; ++i)
            if (cur[row][i] == 'Q')
                return false;
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; --i, --j) 
            if (cur[i][j] == 'Q')
                return false;
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; --i, ++j)
            if (cur[i][j] == 'Q')
                return false;
        return true;
    }
};
