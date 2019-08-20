class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        vector<string> ret;
        unordered_set<string> visited{{s}};
        queue<string> q{{s}};
        bool found = false;
        while (!q.empty()) {
            string temp = q.front();
            q.pop();
            if (isValid(temp)) {
                ret.push_back(temp);
                found = true;
            }
            if (found)
                continue;
            for (int size = temp.size(), i = 0; i < size; ++i) {
                if (temp[i] != '(' && temp[i] != ')')
                    continue;
                string tmp = temp.substr(0, i) + temp.substr(i+1);
                if (!visited.count(tmp)) {
                    q.push(tmp);
                    visited.insert(tmp);
                }
            }
        }
        return ret;
    }
    bool isValid(string t) {
        int cnt = 0, size = t.size();
        for (int i = 0; i < size; ++i) {
            if (t[i] == '(')
                ++cnt;
            else if (t[i] == ')' && --cnt < 0)
                return false;
        }
        return cnt == 0;
    }
};
