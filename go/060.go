func getPermutation(n int, k int) string {
    nums := make([]int, n)
    for i := 1; i <= n; i++ {
        nums[i-1] = i
    }
    var res string
    for len(nums) > 0 {
        divisor := fac(len(nums) - 1)
        idx := (k - 1) / divisor
        res += strconv.Itoa(nums[idx])
        k -= divisor * idx
        temp := nums[:idx]
        temp = append(temp, nums[idx+1:]...)
        nums = temp
    }
    return res
}

func fac(k int) int {
    ret := 1
    if k == 0 || k == 1 {
        return ret
    }
    for ; k > 1; k-- {
        ret *= k
    }
    return ret
}
