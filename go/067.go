func addBinary(a string, b string) string {
    var ret []byte
    i, j := len(a)-1, len(b)-1
    var carry byte
    for i >= 0 && j >= 0 {
        temp := a[i] - '0' + b[j] - '0' + carry
        ret = append(ret, temp%2+'0')
        carry = temp / 2
        i--
        j--
    }
    for i >= 0 {
        temp := a[i] - '0' + carry
        ret = append(ret, temp%2+'0')
        carry = temp / 2
        i--
    }
    for j >= 0 {
        temp := b[j] - '0' + carry
        ret = append(ret, temp%2+'0')
        carry = temp / 2
        j--
    }
    if carry > 0 {
        ret = append(ret, '1')
    }
    size := len(ret)
    for i := 0; i < size/2; i++ {
        ret[i], ret[size-i-1] = ret[size-i-1], ret[i]
    }
    return string(ret)
}

