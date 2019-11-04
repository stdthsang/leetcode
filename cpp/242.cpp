class Solution {
public:
    bool isAnagram(string s, string t) {
        int size1 = s.size(), size2 = t.size();
        if (size1 != size2)
            return false;
        int dict[26] = {0};
        for (auto c : s)
            ++dict[c-'a'];
        for (auto c : t)
            --dict[c-'a'];
        for (int i = 0; i < 26; ++i)
            if (dict[i] != 0)
                return false;
        return true;
    }
};
