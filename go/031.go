func nextPermutation(nums []int) {
    size, idx := len(nums), -1
    for i := size - 1; i > 0; i-- {
        if nums[i] > nums[i-1] {
            idx = i - 1
            break
        }
    }
    if idx == -1 {
        sort.Ints(nums)
        return
    }
    for i := size - 1; i > idx; i-- {
        if nums[i] > nums[idx] {
            nums[i], nums[idx] = nums[idx], nums[i]
            sort.Ints(nums[idx+1:])
            return
        }
    }
}
