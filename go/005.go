func longestPalindrome(s string) string {
    size, ret := len(s), ""
    for i := 0; i < size; i++ {
        s1, s2 := palindrome(s, i, i, size), palindrome(s, i, i+1, size)
        if len(ret) < len(s1) {
            ret = s1
        }
        if len(ret) < len(s2) {
            ret = s2
        }
    }
    return ret
}

func palindrome(s string, i, j, size int) string {
    for i >= 0 && j < size && s[i] == s[j] {
        i--
        j++
    }
    return s[i+1 : j]
}
