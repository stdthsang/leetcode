class Solution {
public:
    int reverse(int x) {
        int ret = 0;
        while (x != 0) {
            if (abs(ret) > INT_MAX / 10 )
                return 0;
            ret = ret * 10 + x % 10;
            x /= 10;
        }
        return ret;
    }
};
