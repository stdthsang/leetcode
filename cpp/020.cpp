class Solution {
public:
    bool isValid(string s) {
        int size = s.size();
        if (size % 2 != 0)
            return false;
        stack<char> stk;
        for (int i = 0; i < size; ++i) {
            if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
                stk.push(s[i]);
                continue;
            }
            if (stk.empty())
                return false;
            switch (s[i]) {
                case ')':
                    if (stk.top() != '(')
                        return false;
                    break;
                case ']':
                    if (stk.top() != '[')
                        return false;
                    break;
                case '}':
                    if (stk.top() != '{')
                        return false;
                    break;
            }
            stk.pop();
        }
        return stk.empty();
    }
};
