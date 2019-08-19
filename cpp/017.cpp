class Solution {
public:
    vector<string> letterCombinations(string digits) {
        if (digits.empty())
            return {};
        int size = digits.size();
        vector<string> ret{""};
        map<char, string> m = {{'2',"abc"},{'3',"def"},{'4',"ghi"},{'5',"jkl"},{'6',"mno"},{'7',"pqrs"},{'8',"tuv"},{'9',"wzxy"}};
        for (int i = 0; i < size; ++i) {
            vector<string> temp;
            int sz = ret.size();
            for (int j = 0; j < sz; ++j) {
                string s = m[digits[i]];
                for (int k = 0, s_sz = s.size(); k < s_sz; ++k) {
                    temp.push_back(ret[j] + s[k]);
                }
            }
            ret = temp;
        }
        return ret;
    }
};
