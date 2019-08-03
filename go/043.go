func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    s1, s2 := len(num1), len(num2)
    ret := make([]byte, s1+s2)
    for i := s1 - 1; i >= 0; i-- {
        var carry byte = 0
        for j := s2 - 1; j >= 0; j-- {
            temp := (num1[i]-'0')*(num2[j]-'0') + carry + ret[i+j+1]
            ret[i+j+1] = temp%10
            carry = temp / 10
        }
        ret[i] += carry
    }

    for i := 0; i < s1 + s2; i++ {
        ret[i] += '0'
    }
    for i := 0; i < s1 + s2; i++ {
        if ret[i] != '0' {
            return string(ret[i:])
        }
    }
    return string(ret)
}
