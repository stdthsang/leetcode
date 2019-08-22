class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int cur = -1, size = nums.size();
        for (int i = 0; i < size; ++i)
            if (nums[i] != val)
                nums[++cur] = nums[i];
        return cur + 1;
    }
};
