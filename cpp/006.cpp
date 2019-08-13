class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows <= 1)
            return s;
        vector<string> strs(numRows, "");
        int size = s.size(), cur = 0, step = 1;
        for (int i = 0; i < size; ++i) {
            strs[cur] += s[i];
            cur += step;
            if (cur == 0) 
                step = 1;
            else if (cur == numRows - 1)
                step = -1;
        }
        string ret;
        for (int i = 0; i < numRows; ++i)
            ret += strs[i];
        return ret;
    }
};
