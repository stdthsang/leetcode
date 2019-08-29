class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int low = equal(nums, target);
        if (low == -1)
            return vector<int>{-1,-1};
        int high = great(nums, target);
        return vector<int>{low, high};
    }
    int equal(vector<int>& nums, int target) {
        int low = 0, high = nums.size() - 1, res = -1;
        while (low <= high) {
            int mid = low + ((high - low) >> 1);
            if (nums[mid] == target) {
                res = mid;
                high = mid - 1;
            } else if (nums[mid] > target)
                high = mid - 1;
            else low = mid + 1;
        }
        return res;
    }
    int great(vector<int>& nums, int target) {
        int low = 0, high = nums.size() - 1, res = -1;
        while (low <= high) {
            int mid = low + ((high - low) >> 1);
            if (nums[mid] <= target) {
                res = mid + 1;
                low = mid + 1;
            } else high = mid - 1;
        }
        return res - 1;
    }
};
