func groupAnagrams(strs []string) [][]string {
    var res [][]string
    dict := make(map[string][]string, 0)
    for _, str := range strs {
        temp := str
        strByte := []byte(str)
        sort.SliceStable(strByte, func(i, j int) bool {
            return strByte[i] < strByte[j]
        })
        dict[string(strByte)] = append(dict[string(strByte)], temp)
    }
    for _, val := range dict {
        res = append(res, val)
    }
    return res
}

