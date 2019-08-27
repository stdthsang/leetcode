func minWindow(s string, t string) string {
    s1, count, dict := len(s), len(t), make([]int, 256)
    for idx := range t {
        dict[t[idx]]++
    }
    cnt, temp, res := 0, make([]int, 256), ""
    for i, j := 0, 0; j < s1; j++ {
        temp[s[j]]++
        if temp[s[j]] <= dict[s[j]] {
            cnt++
        }
        for cnt == count {
            cur := s[i : j+1]
            if res == "" {
                res = cur
            } else if len(cur) < len(res) {
                res = cur
            }

            temp[s[i]]--
            if dict[s[i]] > 0 && temp[s[i]] < dict[s[i]] {
                cnt--
            }
            i++
        }
    }
    return res
}

