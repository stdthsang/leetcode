func longestValidParentheses(s string) int {
    size, ret := len(s), 0
    var stk []int
    var par []byte
    for i := 0; i < size; i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stk = append(stk, i)
            par = append(par, s[i])
            continue
        }
        isMatch := false
        if len(par) > 0 {
            switch s[i] {
            case ')':
                if par[len(par)-1] == '(' {
                    isMatch = true
                }
            case ']':
                if par[len(par)-1] == '[' {
                    isMatch = true
                }
            case '}':
                if par[len(par)-1] == '{' {
                    isMatch = true
                }
            }
        }
        if isMatch {
            stk = stk[:len(stk)-1]
            par = par[:len(par)-1]
        } else {
            stk = append(stk, i)
            par = append(par, s[i])
        }
    }

    last := len(s)
    for len(stk) != 0 {
        curSize := len(stk)
        ret = max(ret, last-stk[curSize-1]-1)
        last = stk[curSize-1]
        stk = stk[:curSize-1]
    }
    ret = max(ret, last)
    return ret
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
