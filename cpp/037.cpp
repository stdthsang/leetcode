class Solution {
public:
    void solveSudoku(vector<vector<char>>& board) {
        help(board, 0, 0);
    }
    bool help(vector<vector<char>>& board, int row, int col) {
        if (col == board[0].size())
            return help(board, row + 1, 0);
        if (row == board.size())
            return true;
        if (board[row][col] != '.')
            return help(board, row, col + 1);
        int m = board.size(), n = board[0].size();
        for (int i = 1; i < 10; ++i) {
            char c = i + '0';
            if (valid(board, row, col, c)) {
                board[row][col] = c;
                if (help(board, row, col + 1))
                    return true;
                board[row][col] = '.';
            }
        }
        return false;
    }
    bool valid(vector<vector<char>>& board, int row, int col, char c) {
        for (int i = 0; i < board.size(); ++i)
            if (board[i][col] == c)
                return false;
        for (int i = 0; i < board[0].size(); ++i)
            if (board[row][i] == c)
                return false;
        for (int i = row / 3 * 3; i < row / 3 * 3 + 3; ++i) {
            for (int j = col / 3 * 3; j < col / 3 * 3 + 3; ++j) {
                if (board[i][j] == c)
                    return false;
            }
        }
        return true;
    }
};
