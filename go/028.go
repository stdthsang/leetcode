func strStr(haystack string, needle string) int {
    s1, s2 := len(haystack), len(needle)
    for i := 0; i <= s1-s2; i++ {
        if haystack[i:i+s2] == needle {
            return i
        }
    }
    return -1
}
