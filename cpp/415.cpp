class Solution {
public:
    string addStrings(string num1, string num2) {
        int s1 = num1.size() - 1, s2 = num2.size() - 1, carry = 0;
        string ret;
        for (; s1 >= 0 && s2 >= 0; --s1,--s2) {
            int temp = num1[s1] - '0' + num2[s2] - '0' + carry;
            ret.push_back(temp%10+'0');
            carry = temp / 10;
        }
        for (; s1 >= 0; --s1) {
            int temp = num1[s1] - '0' + carry;
            ret.push_back(temp%10+'0');
            carry = temp / 10;
        }
        for (; s2 >= 0; --s2) {
            int temp = num2[s2] - '0' + carry;
            ret.push_back(temp%10+'0');
            carry = temp / 10;
        }
        if (carry)
            ret.push_back('1');
        int size = ret.size();
        for (int i = 0; i < size / 2; ++i)
            swap(ret[i], ret[size-i-1]);
        return ret;
    }

};
