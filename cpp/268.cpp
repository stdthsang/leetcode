class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int size = nums.size();
        for (int i = 0; i < size; ++i) {
            while (nums[i] < size && nums[nums[i]] != nums[i])
                swap(nums[nums[i]], nums[i]);
        }
        for (int i = 0; i < size; ++i)
            if (nums[i] != i)
                return i;
        return size;
    }
};
