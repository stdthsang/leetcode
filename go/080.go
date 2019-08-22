func removeDuplicates(nums []int) int {
    cur, size := 0, len(nums)
    for i := 0; i < size; i++ {
        if cur < 2 || nums[i] > nums[cur-2] {
            nums[cur] = nums[i]
            cur++
        }
    }
    return cur
}
