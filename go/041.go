func firstMissingPositive(nums []int) int {
    for i, size := 0, len(nums); i < size; i++ {
        for nums[i] > 0 && nums[i] <= size && nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    size := len(nums)
    for i := 0; i < size; i++ {
        if i+1 != nums[i] {
            return i + 1
        }
    }
    return size + 1
}
