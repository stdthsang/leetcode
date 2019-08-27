class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> out;
        help(nums, 0, out);
        return out;
    }
    
    void help(vector<int>& nums, int start, vector<vector<int>> &out) {
        if (start == nums.size()) {
            out.push_back(nums);
            return;
        }
        int size = nums.size();
        for (int i = start; i < size; ++i) {
            swap(nums[i], nums[start]);
            help(nums, start + 1, out);
            swap(nums[i], nums[start]);
        }
    }
};
