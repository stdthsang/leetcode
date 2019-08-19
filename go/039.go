func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var ret [][]int
    var cur []int
    help(candidates, target, 0, cur, &ret)
    return ret
}

func help(candidates []int, target int, start int, cur []int, out *[][]int) {
    if target == 0 {
        temp := make([]int, len(cur))
        copy(temp, cur)
        *out = append(*out, temp)
        return
    }
    size := len(candidates)
    for i := start; i < size && candidates[i] <= target; i++ {
        cur = append(cur, candidates[i])
        help(candidates, target-candidates[i], i, cur, out)
        cur = cur[:len(cur)-1]
    }
}
