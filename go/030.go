func findSubstring(s string, words []string) []int {
    size, wordCnt, res := len(s), len(words), []int{}
    if size == 0 || wordCnt == 0 {
        return res
    }
    dict, wz := map[string]int{}, len(words[0])
    for _, word := range words {
        dict[word]++
    }
    for i := 0; i < wz; i++ {
        left, cnt, tempDict := i, 0, map[string]int{}
        for j := left; j <= size-wz; j += wz {
            cur := s[j : j+wz]
            if _, ok := dict[cur]; !ok {
                tempDict, cnt, left = map[string]int{}, 0, j+wz
                continue
            }
            tempDict[cur]++
            for tempDict[cur] > dict[cur] {
                t := s[left : left+wz]
                tempDict[t]--
                cnt--
                left += wz
            }
            cnt++
            if cnt == wordCnt {
                res = append(res, left)
                t := s[left : left+wz]
                left += wz
                tempDict[t]--
                cnt--
            }
        }
    }
    return res
}
