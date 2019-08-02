class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        int ret = 0, size = nums.size(), sum = 0;
        for (int i = 0, j = 0; j < size; ) {
            sum += nums[j];
            if (sum == k) {
                ++ret;
                ++j;
            }
            else if (sum < k)
                ++j;
            else
                sum -= nums[i++];
        }
        return ret;
    }
};
