class Solution {
public:
    int divide(int dividend, int divisor) {
        if (divisor == 0 || divisor == -1 && dividend == INT_MIN)
            return INT_MAX;
        int sign = ((dividend > 0) ^ (divisor > 0)) ? -1 : 1;
        long long divAbs = abs((long long)dividend), diAbs = abs((long long)divisor);
        if (diAbs == 1)
            return divAbs * sign;
        int ret = 0;
        while (divAbs >= diAbs) {
            long long temp = diAbs;
            int shift = 1;
            while (divAbs >= (temp << 1)) {
                temp <<= 1;
                shift <<= 1;
            }
            ret += shift;
            divAbs -= temp;
        }
        return ret * sign;
    }
};
