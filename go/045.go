func jump(nums []int) int {
    start, end, res, size := 0, 0, 0, len(nums)
    for start < size-1 {
        res++
        next := 0
        for i := start; i <= end; i++ {
            next = max(next, i+nums[i])
            if next >= size-1 {
                return res
            }
        }
        start, end = end+1, next
    }
    return res
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
