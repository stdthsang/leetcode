class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int size = s.size(), ret = 0;
        int exist[128] = {0};
        for (int i = 0, j = 0; i < size; ++i) {
            while (exist[s[i]]) 
                exist[s[j++]] = 0;
            exist[s[i]] = 1;
            ret = max(ret, i - j + 1);
        }
        return ret;
    }
};
