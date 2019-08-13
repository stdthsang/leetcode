func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
    strs, size, cur, step := make([][]byte, numRows), len(s), 0, 1
    for i := 0; i < size; i++ {
        strs[cur] = append(strs[cur], s[i])
        cur += step
        if cur == 0 {
            step = 1
        } else if cur == numRows-1 {
            step = -1
        }
    }
    var ret string
    for i := 0; i < numRows; i++ {
        ret += string(strs[i])
    }
    return ret
}
