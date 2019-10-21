class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        int size = nums.size();
        for (int i = 0; i < size; ++i)
            while (i < size && nums[nums[i]-1] != nums[i])
                swap(nums[i], nums[nums[i]-1]);
        vector<int> ret;
        for (int i = 0; i < size; ++i)
            if (i != nums[i] - 1)
                ret.push_back(i+1);
        return ret;
    }
};
