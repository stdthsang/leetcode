func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    low, high := 1, x
    for low <= high {
        mid := low + ((high - low) >> 1)
        if mid == x/mid {
            return mid
        }
        if mid > x/mid {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return high
}
