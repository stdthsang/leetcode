class Solution {
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        vector<vector<int>> ret;
        vector<int> cur;
        help(candidates, target, 0, cur, ret);
        return ret;
    }
    void help(vector<int>& candidates, int target, int start, vector<int> cur, vector<vector<int>>& out) {
        if (target == 0) {
            out.push_back(cur);
            return;
        }
        int size = candidates.size();
        for (int i = start; i < size && candidates[i] <= target; ++i) {
            cur.push_back(candidates[i]);
            help(candidates, target-candidates[i], i, cur, out);
            cur.pop_back();
        }
    }
};
