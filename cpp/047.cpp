class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> res;
        permute(nums, 0, res);
        return res;
    }

    void permute(vector<int> nums, int start, vector<vector<int>>& res) {
        if (start == nums.size()) {
            res.push_back(nums);
            return;
        }
        int size = nums.size();
        for (int i = start; i < size; ++i) {
            if (i != start && nums[i] == nums[start]) 
                continue;
            swap(nums[i], nums[start]);
            permute(nums, start + 1, res);
        }
    }
};
