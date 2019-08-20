class Solution {
public:
    int longestValidParentheses(string s) {
        int size = s.size();
        stack<int> stk, par;
        for (int i = 0; i < size; ++i) {
            bool isMath = false;
            if (!par.empty()) {
                if (par.top() == '(' && s[i] == ')')
                    isMath = true;
                if (par.top() == '[' && s[i] == ']')
                    isMath = true;
                if (par.top() == '{' && s[i] == '}')
                    isMath = true;
            }
            if (isMath) {
                stk.pop();
                par.pop();
            } else {
                stk.push(i);
                par.push(s[i]);
            }
        }
        int ret = 0, last = size;
        while (!stk.empty()) {
            ret = max(ret, last-stk.top()-1);
            last = stk.top();
            stk.pop();
        }
        ret = max(ret, last);
        return ret;
    }
};
