func maxSubArray(nums []int) int {
    size := len(nums)
    if size == 0 {
        return 0
    }
    res, sum := nums[0], nums[0]
    for i := 1; i < size; i++ {
        if sum >= 0 {
            sum += nums[i]
        } else {
            sum = nums[i]
        }
        res = max(res, sum)
    }
    return res
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
