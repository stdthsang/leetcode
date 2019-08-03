class Solution {
public:
    string addBinary(string a, string b) {
        string ret;
        int i = a.size() - 1, j = b.size() - 1, carry = 0;
        while (i >= 0 && j >= 0) {
            int temp = a[i--] - '0' + b[j--] - '0' + carry;
            ret.push_back(temp%2+'0');
            carry = temp / 2;
        }
        while (i >= 0) {
            int temp = a[i--] - '0' + carry;
            ret.push_back(temp%2+'0');
            carry = temp / 2;
        }
        while (j >= 0) {
            int temp = b[j--] - '0' + carry;
            ret.push_back(temp%2+'0');
            carry = temp / 2;
        }
        if (carry)
            ret.push_back('1');
        int size = ret.size();
        for (int i = 0; i < size / 2; ++i)
            swap(ret[i], ret[size-i-1]);
        return ret;
    }

};
