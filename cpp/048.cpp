class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        if (matrix.empty() || matrix[0].empty())
            return;
        int row = matrix.size();
        for (int i = 0; i < row; ++i) {
            for (int j = i + 1; j < row; ++j) {
                swap(matrix[i][j], matrix[j][i]);
            }
        }
        for (int i = 0; i < row / 2; ++i) {
            for (int j = 0; j < row; ++j) {
                swap(matrix[j][i], matrix[j][row-i-1]);
            }
        }
    }
};
