func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var ret [][]int
    size := len(nums)
    for i := 0; i < size-2; {
        for j, k := i+1, size-1; j < k; {
            if nums[j]+nums[k]+nums[i] == 0 {
                ret = append(ret, []int{nums[i], nums[j], nums[k]})
                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                for j < k && nums[k] == nums[k-1] {
                    k--
                }
                j++
                k--
            } else if nums[i]+nums[j]+nums[k] < 0 {
                j++
            } else {
                k--
            }
        }
        for i < size-2 && nums[i] == nums[i+1] {
            i++
        }
        i++
    }
    return ret
}
