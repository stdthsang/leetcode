class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        int size = nums.size(), ret = 0, sum = 0;
        unordered_map<int, int> m{{0,1}};
        for (int i = 0; i < size; ++i) {
            sum += nums[i];
            ret += m[sum-k];
            ++m[sum];
        }
        return ret;
    }
};
