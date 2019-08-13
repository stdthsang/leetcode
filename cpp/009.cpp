class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0 || x % 10 == 0 && x != 0)
            return false;
        int x1 = x, y = 0;
        while (x1 != 0) {
            if (y > INT_MAX / 10)
                return false;
            y = y * 10 + x1 % 10;
            x1 /= 10;
        }
        return x == y;
    }
};
