func canJump(nums []int) bool {
    start, end, size := 0, 0, len(nums)
    for start < size {
        next := 0
        for i := start; i <= end; i++ {
            next = max(next, i+nums[i])
            if next >= size-1 {
                return true
            }
        }
        if next <= end {
            return false
        }
        start, end = end+1, next
    }
    return true
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
