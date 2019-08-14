class Solution {
public:
    int maxArea(vector<int>& height) {
        int ret = 0;
        for (int i = 0, j = height.size() - 1; i < j;) {
            ret = max(ret, (j - i) * (height[i] < height[j] ? height[i] : height[j]));
            height[i] < height[j] ? ++i : --j;
        }
        return ret;
    }
};
