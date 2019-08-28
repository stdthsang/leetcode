func search(nums []int, target int) bool {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + ((high - low) >> 1)
        if nums[mid] == target {
            return true
        }
        if nums[mid] > nums[low] {
            if nums[mid] > target && target >= nums[low] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else if nums[mid] < nums[low] {
            if target > nums[mid] && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        } else {
            low++
        }
    }
    return false
}
