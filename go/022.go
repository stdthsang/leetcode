func generateParenthesis(n int) []string {
    var ret []string
    var cur string
    help(n, n, cur, &ret)
    return ret
}

func help(left, right int, cur string, out *[]string) {
    if left > right {
        return
    }
    if left == 0 && right == 0 {
        *out = append(*out, cur)
        return
    }
    if left > 0 {
        help(left-1, right, cur+"(", out)
    }
    if right > 0 {
        help(left, right-1, cur+")", out)
    }
}
