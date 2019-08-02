class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        int size = nums.size();
        sort(nums.begin(), nums.end());
        vector<vector<int>> ret;
        for (int i = 0; i < size - 3; ) {
            for (int j = i + 1; j < size - 2; ) {
                for (int k = j + 1, l = size - 1; k < l; ) {
                    long long sum = nums[i] + nums[j] + nums[k] + nums[l];
                    if (sum == target) {
                        vector<int> temp{nums[i], nums[j], nums[k], nums[l]};
                        ret.push_back(temp);
                        while (k < l && nums[k] == nums[k+1])
                            ++k;
                        while (k < l && nums[l] == nums[l-1])
                            --l;
                        ++k, --l;
                    } else if (sum < target)
                        ++k;
                    else --l;
                }
                while (j < size - 2 && nums[j] == nums[j+1])
                    ++j;
                ++j;
            }
            while (i < size - 3 && nums[i] == nums[i+1])
                ++i;
            ++i;
        }
        return ret;
    }
};
