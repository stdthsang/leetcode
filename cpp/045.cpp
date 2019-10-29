class Solution {
public:
    int jump(vector<int>& nums) {
        int start = 0, end = 0, res = 0, size = nums.size();
        for (; start < size - 1;) {
            ++res;
            int next = 0;
            for (int i = start; i <= end; ++i) {
                next = max(next, i + nums[i]);
                if (next >= size - 1)
                    return res;
            }
            start = end + 1;
            end = next;
        }
        return res;
    }
};
