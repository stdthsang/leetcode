class Solution {
public:
    int myAtoi(string str) {
        int ret = 0, size = str.size(), i = 0, sign = 1;
        while (i < size && isspace(str[i]))
            ++i;
        if (i < size && (str[i] == '+' || str[i] == '-')) {
            sign = str[i] == '+' ? 1 : -1;
            ++i;
        }
        for (; i < size && str[i] >= '0' && str[i] <= '9'; ++i) {
            if (ret > INT_MAX / 10 || (ret == INT_MAX/ 10 && str[i] > '7'))
                return sign == 1 ? INT_MAX : INT_MIN;
            ret = ret * 10 + (str[i] - '0');
        }
        return ret * sign;
    }
};
