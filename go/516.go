func longestPalindromeSubseq(s string) int {
    size, dp := len(s), make([][]int, len(s))
    for i := 0; i < size; i++ {
        dp[i] = make([]int, size)
    }
    for i := size - 1; i >= 0; i-- {
        dp[i][i] = 1
        for j := i + 1; j < size; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    return dp[0][size-1]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
