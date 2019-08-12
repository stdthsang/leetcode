func countSubstrings(s string) int {
    size, ret := len(s), 0
    dp := make([][]bool, size)
    for i := 0; i < size; i++ {
        dp[i] = make([]bool, size)
    }
    for i := size - 1; i >= 0; i-- {
        for j := i; j < size; j++ {
            if (s[i] == s[j]) && (j-i <= 2 || dp[i+1][j-1]) {
                dp[i][j] = true
            }
            if dp[i][j] {
                ret++
            }
        }
    }
    return ret
}
