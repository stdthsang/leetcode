class Solution {
public:
    vector<vector<int>> palindromePairs(vector<string>& words) {
        int size = words.size();
        vector<vector<int>> ret;
        unordered_map<string, int> dict;
        for (int i = 0; i < size; ++i) {
            string key = words[i];
            reverse(key.begin(), key.end());
            dict[key] = i;
        }
        // edge case: if empty string "" exists, find all palindromes to become pairs ("", self)
        if (dict.find("") != dict.end()) {
            for (int i = 0; i < size; ++i) {
                if (i == dict[""])
                    continue;
                if (isPalindrome(words[i]))
                    ret.push_back({dict[""], i});
            }
        }
        for (int i = 0; i < size; ++i) {
            int wl = words[i].size();
            for (int j = 0; j < wl; ++j) {
                string left = words[i].substr(0, j);
                string right = words[i].substr(j, wl - j);
                if (dict.find(left) != dict.end() && isPalindrome(right) && dict[left] != i)
                    ret.push_back({i, dict[left]});
                if (dict.find(right) != dict.end() && isPalindrome(left) && dict[right] != i)
                    ret.push_back({dict[right], i});
            }
        }
        return ret;
    }

    bool isPalindrome(const string s) {
        int size = s.size();
        for (int i = 0; i < size / 2; ++i)
            if (s[i] != s[size-i-1])
                return false;
        return true;
    }
};
