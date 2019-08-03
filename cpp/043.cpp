ass Solution {
public:
    string multiply(string num1, string num2) {
        int s1 = num1.size(), s2 = num2.size();
        vector<int> res(s1+s2, 0);
        for (int i = s1 - 1; i >= 0; --i) {
            int carry = 0;
            for (int j = s2 - 1; j >= 0; --j) {
                int temp = (num1[i] - '0') * (num2[j] - '0') + carry + res[i+j+1];
                res[i+j+1] = temp % 10;
                carry = temp / 10;
            }
            res[i] = carry;
        }
        string ret;
        for (int i = 0; i < s1 + s2; ++i)
            ret.push_back(res[i]+'0');
        for (int i = 0; i < s1 + s2; ++i) 
            if (ret[i] != '0')
                return ret.substr(i);
        return "0";
    }
};
