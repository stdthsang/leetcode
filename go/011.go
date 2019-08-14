func maxArea(height []int) int {
    ret := 0
    for i, j := 0, len(height)-1; i < j; {
        ret = max(ret, (j-i)*(min(height[i], height[j])))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }
    return ret
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
