func romanToInt(s string) int {
    size, ret := len(s), 0
    m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    for i := 0; i < size; i++ {
        if i == 0 || m[s[i]] <= m[s[i-1]] {
            ret += m[s[i]]
        } else {
            ret += m[s[i]] - 2*m[s[i-1]]
        }
    }
    return ret
}
