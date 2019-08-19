func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for len(strs) > 1 {
        var temp []string
        for i := 0; i < len(strs)-1; i += 2 {
            tmp := commonPrefix(strs[i], strs[i+1])
            if tmp == "" {
                return ""
            }
            temp = append(temp, tmp)
        }
        if len(strs)%2 != 0 {
            temp = append(temp, strs[len(strs)-1])
        }
        strs = temp
    }
    return strs[0]
}

func commonPrefix(str1, str2 string) string {
    size, s2 := len(str1), len(str2)
    if size > s2 {
        size = s2
    }
    var ret []byte
    for i := 0; i < size; i++ {
        if str1[i] != str2[i] {
            return string(ret)
        }
        ret = append(ret, str1[i])
    }
    return string(ret)
}
