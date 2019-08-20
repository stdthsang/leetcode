func removeInvalidParentheses(s string) []string {
    var ret []string
    visited, queue, found := map[string]bool{s: true}, []string{s}, false
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        if isValid(cur) {
            ret = append(ret, cur)
            found = true
        }
        if found {
            continue
        }
        for idx, c := range cur {
            if c != '(' && c != ')' {
                continue
            }
            temp := cur[:idx] + cur[idx+1:]
            if _, ok := visited[temp]; !ok {
                visited[temp] = true
                queue = append(queue, temp)
            }
        }
    }
    return ret
}

func isValid(str string) bool {
    cnt := 0
    for _, c := range str {
        if c == '(' {
            cnt++
        } else if c == ')' {
            cnt--
            if cnt < 0 {
                return false
            }
        }
    }
    return cnt == 0
}

