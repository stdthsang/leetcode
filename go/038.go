func countAndSay(n int) string {
    if n < 1 {
        return ""
    }
    ret, m := "1", map[int]string{1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}
    for ; n > 1; n-- {
        size, cur := len(ret), ""
        for i, j := 0, 0; j < size; {
            for j < size && ret[j] == ret[i] {
                j++
            }
            cur += m[j-i] + m[int(ret[i]-'0')]
            if j == size {
                break
            }
            i = j
        }
        ret = cur
    }
    return ret
}
