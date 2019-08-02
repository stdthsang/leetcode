class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int>> ret;
        int size = nums.size();
        for (int i = 0; i < size - 2; ) {
            for (int j = i + 1, k = size - 1; j < k; ) {
                long long sum = nums[i] + nums[j] + nums[k];
                if (sum == 0) {
                    vector<int> temp{nums[i], nums[j], nums[k]};
                    ret.push_back(temp);
                    while (j < k && nums[j] == nums[j+1])
                        ++j;
                    while (j < k && nums[k] == nums[k-1])
                        --k;
                    ++j;--k;
                } else if (sum < 0)
                    ++j;
                else --k;
            }
            while (i < size - 2 && nums[i] == nums[i+1])
                ++i;
            ++i;
        }
        return ret;
    }
};
