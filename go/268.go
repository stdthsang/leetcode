func missingNumber(nums []int) int {
    size := len(nums)
    for i := 0; i < size; i++ {
        for nums[i] < size && nums[nums[i]] != nums[i] {
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        }
    }
    for i := 0; i < size; i++ {
        if i != nums[i] {
            return i
        }
    }
    return size
}
