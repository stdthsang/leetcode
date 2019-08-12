func shortestPalindrome(s string) string {
    i, n := 0, len(s)
    for j := n - 1; j >= 0; j-- {
        if s[i] == s[j] {
            i++
        }
    }
    if i == n {
        return s;
    }
    var rem []byte
    for j := n - 1; j >= i; j-- {
        rem = append(rem, s[j])
    }
    return string(rem) + shortestPalindrome(s[:i]) + s[i:]
}
