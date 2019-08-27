class Solution {
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        vector<int> ret;
        if (s.empty() || words.empty())
            return ret;
        int size = s.size(), wz = words[0].size(), cnt = words.size();
        map<string, int> dict;
        for (auto str : words)
            ++dict[str];
        for (int i = 0; i < wz; ++i) {
            int left = i, matchCnt = 0;
            map<string, int> temp;
            for (int j = i; j <= size - wz; j += wz) {
                string cur = s.substr(j, wz);
                if (dict.find(cur) == dict.end()) {
                    temp.clear();
                    left = j + wz;
                    matchCnt = 0;
                    continue;
                }
                ++temp[cur];
                while (temp[cur] > dict[cur]) {
                    string t = s.substr(left, wz);
                    --temp[t];
                    --matchCnt;
                    left += wz;
                }
                if (++matchCnt == cnt) {
                    ret.push_back(left);
                    string t = s.substr(left, wz);
                    --temp[t];
                    --matchCnt;
                    left += wz;
                }
            }
        }
        return ret;
    }
};
