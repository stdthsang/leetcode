class Solution {
public:
    int findDuplicate(vector<int>& nums) {
        int low = 1, high = nums.size() - 1;
        while (low <= high) {
            int mid = low + ((high - low) >> 1), count = 0;
            for (auto n : nums)
                if (n <= mid)
                    ++count;
            if (count <= mid)
                low = mid + 1;
            else high = mid - 1;
        }
        return low;
    }
};
