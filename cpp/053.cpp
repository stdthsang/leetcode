class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int size = nums.size();
        if (size == 0)
            return 0;
        int res = nums[0], sum = nums[0];
        for (int i = 1; i < size; ++i) {
            if (sum >= 0)
                sum += nums[i];
            else sum = nums[i];
            res = max(res, sum);
        }
        return res;
    }
};
