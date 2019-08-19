class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        vector<int> ret;
        for (int i = 0, j = numbers.size() - 1; i < j; ) {
            int sum = numbers[i] + numbers[j];
            if (sum == target) {
                ret.push_back(i+1);
                ret.push_back(j+1);
                return ret;
            } else if (sum < target)
                ++i;
            else --j;
        }
        return ret;
    }
};

