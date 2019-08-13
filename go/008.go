func myAtoi(str string) int {
    size, i, sign := len(str), 0, 1
    for i < size && str[i] == ' ' {
        i++
    }
    if i < size && (str[i] == '+' || str[i] == '-') {
        if str[i] == '-' {
            sign = -1
        }
        i++
    }
    var ret int32
    for ; i < size && str[i] >= '0' && str[i] <= '9'; i++ {
        if ret > math.MaxInt32/10 || ret == math.MaxInt32/10 && str[i] > '7' {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        ret = ret*10 + int32(str[i] - '0')
    }
    return sign * int(ret)
}
