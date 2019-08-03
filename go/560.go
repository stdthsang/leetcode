subarraySum(nums []int, k int) int {
    ret, m, sum := 0, map[int]int{0: 1}, 0
    for _, num := range nums {
        sum += num
        ret += m[sum-k]
        m[sum]++
    }
    return ret
}
