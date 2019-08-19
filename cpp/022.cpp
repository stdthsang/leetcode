class Solution {
public:
    vector<string> generateParenthesis(int n) {
        vector<string> ret;
        string cur;
        help(n, n, cur, ret);
        return ret;
    }
    void help(int left, int right, string cur, vector<string> &out) {
        if (right < left)
            return;
        if (left == 0 && right == 0) {
            out.push_back(cur);
            return;
        }
        if (left > 0)
            help(left - 1, right, cur + "(", out);
        if (right > 0)
            help(left, right - 1, cur + ")", out);
    }
};
