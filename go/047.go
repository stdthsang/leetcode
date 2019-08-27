func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    help(nums, 0, &res)
    return res
}

func help(nums []int, start int, res *[][]int) {
    if start == len(nums) {
        temp := make([]int, len(nums))
        copy(temp, nums)
        *res = append(*res, temp)
        return
    }
    size := len(nums)
    for i := start; i < size; i++ {
        if i != start && nums[i] == nums[start] {
            continue
        }
        nums[i], nums[start] = nums[start], nums[i]
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        help(tmp, start+1, res)
    }
}
