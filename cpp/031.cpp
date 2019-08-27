class Solution {
public:
    void nextPermutation(vector<int>& nums) {
        int size = nums.size(), idx = -1;
        for (int i = size - 1; i >= 1; --i) {
            if (nums[i] > nums[i-1]) {
                idx = i - 1;
                break;
            }
        }
        if (idx == -1) {
            sort(nums.begin(), nums.end());
            return;
        }
        for (int i = size - 1; i > idx; --i) {
            if (nums[i] > nums[idx]) {
                swap(nums[i],nums[idx]);
                sort(nums.begin()+idx+1, nums.end());
                return;
            }
        }
    }
};
