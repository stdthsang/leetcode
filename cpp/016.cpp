class Solution {
public:
    int threeSumClosest(vector<int>& nums, int target) {
        sort(nums.begin(), nums.end());
        int size = nums.size(), closest = INT_MAX, ret = 0;
        for (int i = 0; i < size - 2; ++i) {
            for (int j = i + 1, k = size - 1; j < k; ) {
                int temp = abs(nums[i]+nums[j]+nums[k]-target);
                if (temp == 0)
                    return target;
                if (temp < closest) {
                    closest = temp;
                    ret = nums[i]+nums[j]+nums[k];
                } 
                if (nums[i]+nums[j]+nums[k] < target)
                    ++j;
                else --k;
            }
        }
        return ret;
    }
};
