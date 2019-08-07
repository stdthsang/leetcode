class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int s1 = nums1.size(), s2 = nums2.size();
        int size = s1 + s2;
        if (size % 2 == 0)
            return (static_cast<double>(findKth(nums1, 0, s1 - 1, nums2, 0, s2 - 1, size / 2)) +
                    static_cast<double>(findKth(nums1, 0, s1 - 1, nums2, 0, s2 - 1, size / 2 + 1))) / 2;
        return static_cast<double>(findKth(nums1, 0, s1 - 1, nums2, 0, s2 - 1, size / 2 + 1));
    }
    int findKth(const vector<int>& nums1, int s1, int e1, const vector<int>& nums2, int s2, int e2, int k) {
        int n1 = e1 - s1 + 1, n2 = e2 - s2 + 1;
        if (n1 < n2)
            return findKth(nums2, s2, e2, nums1, s1, e1, k);
        if (n2 <= 0)
            return nums1[s1+k-1];
        if (k == 1)
            return nums1[s1] <= nums2[s2] ? nums1[s1] : nums2[s2];
        int posB = k / 2 > n2 ? n2 : k / 2, posA = k - posB;
        if (nums1[s1+posA-1] < nums2[s2+posB-1])
            return findKth(nums1, s1 + posA, e1, nums2, s2, e2, k - posA);
        return findKth(nums1, s1, e1, nums2, s2 + posB, e2, k - posB);
    }
};
