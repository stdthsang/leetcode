func trap(height []int) int {
    size, ret, mIndex := len(height), 0, 0
    for i := range height {
        if height[i] > height[mIndex] {
            mIndex = i
        }
    }
    cur := 0
    for i := 0; i < mIndex; i++ {
        if cur > height[i] {
            ret += cur - height[i]
        } else {
            cur = height[i]
        }
    }
    cur = 0
    for i := size - 1; i > mIndex; i-- {
        if cur > height[i] {
            ret += cur - height[i]
        } else {
            cur = height[i]
        }
    }
    return ret
}

