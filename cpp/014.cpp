class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.empty())
            return "";
        while (strs.size() > 1) {
            vector<string> temp;
            for (int i = 0; i < strs.size() - 1; i += 2) {
                string tmp = commonPrefix(strs[i], strs[i+1]);
                if (tmp == "")
                    return "";
                temp.push_back(tmp);
            }
            if (strs.size() % 2 != 0)
                temp.push_back(strs[strs.size()-1]);
            strs = temp;
        }
        return strs[0];
    }
    
    string commonPrefix(const string &s1, const string &s2) {
        string ret;
        int size = s1.size(), sz = s2.size();
        if (size > sz)
            size = sz;
        for (int i = 0; i < size; ++i) {
            if (s1[i] != s2[i])
                return ret;
            ret += s1[i];
        }
        return ret;
    }
};
