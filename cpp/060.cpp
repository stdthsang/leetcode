class Solution {
public:
    string getPermutation(int n, int k) {
        vector<int> nums;
        for (int i = 1; i <= n; ++i)
            nums.push_back(i);
        string res;
        while (!nums.empty()) {
            int divisor = fac(nums.size() - 1);
            int idx = (k - 1) / divisor;
            res += nums[idx] + '0';
            k -= divisor * idx;
            nums.erase(nums.begin()+idx);
        }
        return res;
    }

    int fac(int n) {
        if (n == 0 || n == 1)
            return 1;
        int res = 1;
        for (int i = 2; i <= n; ++i)
            res *= i;
        return res;
    }
};
