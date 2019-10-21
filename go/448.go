func findDisappearedNumbers(nums []int) []int {
    size := len(nums)
    for i := 0; i < size; i++ {
        for i < size && nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    var ret []int
    for i := 0; i < size; i++ {
        if i != nums[i]-1 {
            ret = append(ret, i+1)
        }
    }
    return ret
}
