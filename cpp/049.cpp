class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        vector<vector<string>> res;
        map<string, vector<string>> dict;
        for (auto str : strs) {
            string temp = str;
            sort(str.begin(), str.end());
            if (dict.count(str))
                dict[str].push_back(temp);
            else
                dict[str] = vector<string>{temp};
        }
        for (map<string, vector<string>>::iterator iter = dict.begin(); iter != dict.end(); ++iter) {
            res.push_back(iter->second);
        }
        return res;
    }
};
