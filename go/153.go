func findMin(nums []int) int {
    low, high := 0, len(nums)-1
    for low < high {
        mid := low + ((high - low) >> 1)
        if nums[mid] > nums[high] {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return nums[low]
}
