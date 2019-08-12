class Solution {
public:
    int countSubstrings(string s) {
        int size = s.size(), ret = 0;
        vector<vector<bool>> dp(size, vector<bool>(size, false));
        for (int i = size - 1; i >= 0; --i) {
            for (int j = i; j < size; ++j) {
                dp[i][j] = (s[i] == s[j]) && (j - i <= 2 || dp[i+1][j-1]);
                if (dp[i][j])
                    ++ret;
            }
        }
        return ret;
    }
};
