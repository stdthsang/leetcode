func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    s1, s2 := len(nums1), len(nums2)
    size := s1 + s2
    if size%2 == 0 {
        return (float64(findKth(nums1, 0, s1-1, nums2, 0, s2-1, size/2)) +
            float64(findKth(nums1, 0, s1-1, nums2, 0, s2-1, size/2+1))) / 2
    }
    return float64(findKth(nums1, 0, s1-1, nums2, 0, s2-1, size/2+1))
}

func findKth(nums1 []int, s1, e1 int, nums2 []int, s2, e2 int, k int) int {
    n1, n2 := e1-s1+1, e2-s2+1
    if n1 < n2 {
        return findKth(nums2, s2, e2, nums1, s1, e1, k)
    }
    if n2 <= 0 {
        return nums1[s1+k-1]
    }
    if k == 1 {
        return min(nums1[s1], nums2[s2])
    }
    posB := min(k/2, n2)
    posA := k - posB
    if nums1[s1+posA-1] <= nums2[s2+posB-1] {
        return findKth(nums1, s1+posA, e1, nums2, s2, e2, k-posA)
    }
    return findKth(nums1, s1, e1, nums2, s2+posB, e2, k-posB)
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

