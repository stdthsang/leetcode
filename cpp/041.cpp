class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        for (int i = 0, size = nums.size(); i < size; ++i) {
            while (nums[i] > 0 && nums[i] <= size && nums[nums[i]-1] != nums[i])
                swap(nums[nums[i]-1], nums[i]);
        }
        int size = nums.size();
        for (int i = 0; i < size; i++) {
            if (i + 1 != nums[i])
                return i + 1;
        }
        return size + 1;
    }
};
