func isValid(S string) bool {
    target, size := "abc", len(S)
    if size%3 != 0 {
        return false
    }
    for len(S) >= 3 {
        idx := strings.Index(S, target)
        if idx == -1 {
            return false
        }
        S = S[:idx] + S[idx+3:]
    }
    return len(S) == 0
}

func isValid1(S string) bool {
    var stack []byte
    for _, c := range S {
        if c == 'c' {
            size := len(stack)
            if size < 2 || stack[size-1] != 'b' || stack[size-2] != 'a' {
                return false
            }
            stack = stack[:size-2]
        } else {
            stack = append(stack, byte(c))
        }
    }
    return len(stack) == 0
}
