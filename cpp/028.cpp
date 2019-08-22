class Solution {
public:
    int strStr(string haystack, string needle) {
        int s1 = haystack.size(), s2 = needle.size();
        for (int i = 0; i <= s1 - s2; ++i)
            if (haystack.substr(i, s2) == needle)
                return i;
        return -1;
    }
};
