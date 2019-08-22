class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.empty())
            return 0;
        int cur = 0, size = nums.size();
        for (int i = 1; i < size; ++i) {
            if (nums[cur] != nums[i])
                nums[++cur] = nums[i];
        }
        return cur + 1;
    }
};
