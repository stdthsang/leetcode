class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        vector<vector<bool>> row(9, vector<bool>(9));
        vector<vector<bool>> col(9, vector<bool>(9));
        vector<vector<bool>> cell(9, vector<bool>(9));
        for (int i = 0; i < 9; ++i) {
            for (int j = 0; j < 9; ++j) {
                if (board[i][j] == '.')
                    continue;
                int c = board[i][j] - '1';
                if (row[i][c] || col[c][j] || cell[3 * (i / 3) + j / 3][c])
                    return false;
                row[i][c] = col[c][j] = cell[3 * (i / 3) + j / 3][c] = true;
            }
        }
        return true;
    }
};
