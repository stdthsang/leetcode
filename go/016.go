func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    size, ret, closest := len(nums), 0, math.MaxInt32
    for i := 0; i < size-2; i++ {
        for j, k := i+1, size-1; j < k; {
            temp := nums[i] + nums[j] + nums[k]
            if temp == target {
                return target
            } else if temp < target {
                j++
            } else {
                k--
            }
            newDiff := abs(temp - target)
            if newDiff < closest {
                closest = newDiff
                ret = temp
            }
        }
    }
    return ret
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
