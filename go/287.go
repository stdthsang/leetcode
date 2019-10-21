func findDuplicate(nums []int) int {
    low, high := 1, len(nums)-1
    for low <= high {
        mid, count := low+((high-low)>>1), 0
        for _, num := range nums {
            if num <= mid {
                count++
            }
        }
        if count <= mid {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return low
}
