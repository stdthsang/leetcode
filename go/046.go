func permute(nums []int) [][]int {
    var out [][]int
    help(nums, 0, &out)
    return out
}

func help(nums []int, start int, out *[][]int) {
    if start == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        *out = append(*out, temp)
        return
    }
    size := len(nums)
    for i := start; i < size; i++ {
        nums[i], nums[start] = nums[start], nums[i]
        help(nums, start+1, out)
        nums[i], nums[start] = nums[start], nums[i]
    }
}

