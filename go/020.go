func isValid(s string) bool {
    size := len(s)
    var stk []byte
    for i := 0; i < size; i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stk = append(stk, s[i])
            continue
        }
        if len(stk) == 0 {
            return false
        }
        switch s[i] {
        case ')':
            if stk[len(stk)-1] != '(' {
                return false
            }
        case ']':
            if stk[len(stk)-1] != '[' {
                return false
            }
        case '}':
            if stk[len(stk)-1] != '{' {
                return false
            }
        }
        stk = stk[:len(stk)-1]
    }
    return len(stk) == 0
}
