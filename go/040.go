func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var cur []int
    var ret [][]int
    helper(candidates, target, 0, cur, &ret)
    return ret
}

func helper(candidates []int, target int, start int, cur []int, ret *[][]int) {
    if target == 0 {
        var temp []int
        for _, v := range cur {
            temp = append(temp, v)
        }
        *ret = append(*ret, temp)
        return
    }
    for i, size := start, len(candidates); i < size && target >= candidates[i]; i++ {
        cur = append(cur, candidates[i])
        helper(candidates, target-candidates[i], i+1, cur, ret)
        cur = cur[:len(cur)-1]
        for i < size-1 && candidates[i] == candidates[i+1] {
            i++
        }
    }
}

