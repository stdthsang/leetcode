func removeElement(nums []int, val int) int {
    cur, size := -1, len(nums)
    for i := 0; i < size; i++ {
        if nums[i] != val {
            cur++
            nums[cur] = nums[i]
        }
    }
    return cur + 1
}
