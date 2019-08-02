func twoSum_1(nums []int, target int) []int {
    size, ret, m := len(nums), []int{}, map[int]int{}
    for i := 0; i < size; i++ {
        if v, ok := m[target-nums[i]]; ok {
            ret = append(ret, v, i)
            return ret
        }
        m[nums[i]] = i
    }
    return ret
}

func twoSum(nums []int, target int) []int {
    size, ret := len(nums), []int{}
    var ns nodes;
    for i, val := range nums {
        ns = append(ns, node{i, val})
    }
    sort.SliceStable(ns, func(i, j int) bool {
        return ns[i].val < ns[j].val
    })
    //sort.Sort(ns)
    for i, j := 0, size - 1; i < j; {
        if ns[i].val + ns[j].val == target {
            if ns[i].index < ns[j].index {
                ret = append(ret, ns[i].index, ns[j].index)
            } else {
                ret = append(ret, ns[j].index, ns[i].index)
            }
            return ret
        }
        if ns[i].val + ns[j].index < target {
            i++
        } else {
            j--
        }
    }
    return ret
}

type node struct {
    index int
    val int
}

type nodes []node

func (n nodes) Less(i, j int) bool {
    return n[i].val < n[j].val
}

func (n nodes) Len() int {
    return len(nodes)
}

func (n nodes) Swap(i, j int) {
    n[i], n[j] = n[j], n[i]
}
