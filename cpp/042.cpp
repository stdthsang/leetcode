class Solution {
public:
    int trap(vector<int>& height) {
        int size = height.size(), ret = 0, mIndex = 0;
        for (int i = 0; i < size; ++i)
            if (height[i] > height[mIndex])
                mIndex = i;
        
        for (int i = 0, cur = 0; i < mIndex; ++i) {
            if (height[i] > cur)
                cur = height[i];
            else
                ret += cur - height[i];
        }
        for (int i = size - 1, cur = 0; i > mIndex; --i) {
            if (height[i] > cur)
                cur = height[i];
            else
                ret += cur - height[i];
        }
        return ret;
    }
};
