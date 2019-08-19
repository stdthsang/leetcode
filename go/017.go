func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    m := map[byte][]byte{'2': []byte("abc"), '3': []byte("def"), '4': []byte("ghi"),
        '5': []byte("jkl"), '6': []byte("mno"), '7': []byte("pqrs"), '8': []byte("tuv"), '9': []byte("wxzy")}
    ret, size := []string{""}, len(digits)
    for i := 0; i < size; i++ {
        temp, sz := []string{}, len(ret)
        for j := 0; j < sz; j++ {
            ssss, ok := m[digits[i]]
            if !ok {
                return nil
            }
            bytes := []byte(ret[j])
            for _, c := range ssss {
                bytes = append(bytes, c)
                temp = append(temp, string(bytes))
                bytes = bytes[:len(bytes)-1]
            }
        }
        ret = temp
    }
    return ret
}
