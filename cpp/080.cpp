class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int cur = 0, size = nums.size();
        for (int i = 0; i < size; ++i) {
            if (cur < 2 || nums[i] > nums[cur-2])
                nums[cur++] = nums[i];
        }
        return cur;
    }
};
