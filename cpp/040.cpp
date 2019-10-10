class Solution {
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        vector<vector<int>> ret;
        vector<int> cur;
        helper(candidates, target, 0, cur, ret);
        return ret;
    }
    void helper(vector<int>& candidates, int target, int start, vector<int> cur, vector<vector<int>> &ret) {
        if (target == 0) {
            ret.push_back(cur);
            return;
        }
        int size = candidates.size();
        for (int i = start; i < size && candidates[i] <= target; ++i) {
            cur.push_back(candidates[i]);
            helper(candidates, target-candidates[i], i + 1, cur, ret);
            cur.pop_back();
            while (i < size - 1 && candidates[i] == candidates[i+1])
                ++i;
        }
    }
};
