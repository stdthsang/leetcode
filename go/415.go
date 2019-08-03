func addStrings(num1 string, num2 string) string {
    var ret []byte
    var carry byte
    i, j := len(num1)-1, len(num2)-1
    for i >= 0 && j >= 0 {
        temp := num1[i] - '0' + num2[j] - '0' + carry
        ret = append(ret, temp%10+'0')
        carry = temp / 10
        i--
        j--
    }
    for i >= 0 {
        temp := num1[i] - '0' + carry
        ret = append(ret, temp%10+'0')
        carry = temp / 10
        i--
    }
    for j >= 0 {
        temp := num2[j] - '0' + carry
        ret = append(ret, temp%10+'0')
        carry = temp / 10
        j--
    }
    if carry != 0 {
        ret = append(ret, '1')
    }
    size := len(ret)
    for i := 0; i < size/2; i++ {
        ret[i], ret[size-i-1] = ret[size-i-1], ret[i]
    }
    return string(ret)
}
