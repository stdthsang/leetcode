// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int low = 1, high = n, res = -1;
        while (low <= high) {
            int mid = low + ((high - low) >> 1);
            if (isBadVersion(mid)) {
                res = mid;
                high = mid - 1;
            } else low = mid + 1;
        }
        return res;
    }
};
