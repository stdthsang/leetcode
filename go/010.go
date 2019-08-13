func isMatch(s string, p string) bool {
    s1, s2 := len(s), len(p)
    dp := make([][]bool, s1+1)
    for i := 0; i <= s1; i++ {
        dp[i] = make([]bool, s2+1)
    }
    dp[0][0] = true
    for i := 1; i <= s2; i++ {
        dp[0][i] = i > 1 && p[i-1] == '*' && dp[0][i-2]
    }
    for i := 1; i <= s1; i++ {
        for j := 1; j <= s2; j++ {
            if p[j-1] != '*' {
                dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
            } else {
                dp[i][j] = j > 1 && (dp[i][j-2] || ((s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j]))
            }
        }
    }
    return dp[s1][s2]
}
