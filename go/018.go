func fourSum(nums []int, target int) [][]int {
    size, ret := len(nums), [][]int{}
    sort.Ints(nums)
    for i := 0; i < size-3; {
        for j := i + 1; j < size-2; {
            for k, l := j+1, size-1; k < l; {
                sum := nums[i] + nums[j] + nums[k] + nums[l];
                if sum == target {
                    ret = append(ret, []int{nums[i], nums[j], nums[k], nums[l]})
                    for k < l && nums[k] == nums[k+1] {
                        k++
                    }
                    for k < l && nums[l] == nums[l-1] {
                        l--
                    }
                    k++
                    l--
                } else if sum < target {
                    k++
                } else {
                    l--
                }
            }
            for j < size-2 && nums[j] == nums[j+1] {
                j++
            }
            j++
        }
        for i < size-3 && nums[i] == nums[i+1] {
            i++
        }
        i++
    }
    return ret
}
