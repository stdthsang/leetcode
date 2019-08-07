func lengthOfLongestSubstring(s string) int {
    size, exist, ret := len(s), make([]bool, 128), 0
    for i, j := 0, 0; i < size; i++ {
        for exist[s[i]] {
            exist[s[j]] = false
            j++
        }
        exist[s[i]] = true
        ret = max(ret, i-j+1)
    }
    return ret
}

func max(i, j int) int {
    if i < j {
        return j
    }
    return i
}
