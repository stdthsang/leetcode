func removeDuplicates(nums []int) int {
    cur, size := 0, len(nums)
    if size == 0 {
        return 0
    }
    for i := 1; i < size; i++ {
        if nums[cur] != nums[i] {
            cur++
            nums[cur] = nums[i]
        }
    }
    return cur + 1
}
