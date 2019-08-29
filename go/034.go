func searchRange(nums []int, target int) []int {
    low := equal(nums, target)
    if low == -1 {
        return []int{-1, -1}
    }
    high := great(nums, target)
    return []int{low, high}
}

func equal(nums []int, target int) int {
    low, high, res := 0, len(nums)-1, -1
    for low <= high {
        mid := low + ((high - low) >> 1)
        if nums[mid] == target {
            res = mid
            high = mid - 1
        } else if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return res
}

func great(nums []int, target int) int {
    low, high, res := 0, len(nums)-1, -1
    for low <= high {
        mid := low + ((high - low) >> 1)
        if nums[mid] <= target {
            res, low = mid+1, mid+1
        } else {
            high = mid - 1
        }
    }
    return res - 1
}
