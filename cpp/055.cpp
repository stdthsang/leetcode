class Solution {
public:
    bool canJump(vector<int>& nums) {
        int start = 0, end = 0, size= nums.size();
        for (; start < size; ) {
            int next = 0;
            for (int i = start; i <= end; ++i) {
                next = max(next, i + nums[i]);
                if (next >= size - 1)
                    return true;
            }
            if (next <= end)
                return false;
            start = end + 1;
            end = next;
        }
        return true;
    }
};
