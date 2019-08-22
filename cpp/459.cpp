class Solution {
public:
    bool repeatedSubstringPattern(string s) {
        int size = s.size();
        for (int i = 0; i < size / 2; ++i) {
            string sub = s.substr(0, i + 1);
            int subSize = sub.size();
            if (size % subSize != 0)
                continue;
            int j = subSize;
            for (; j < size; j += subSize) {
                if (sub != s.substr(j, subSize))
                    break;
            }
            if (j == size && s.substr(j-subSize, subSize) == sub)
                return true;
        }
        return false;
    }
};
