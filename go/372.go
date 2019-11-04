func superPow(a int, b []int) int {
    size, res := len(b), 1
    for i := 0; i < size; i++ {
        res = pow(res, 10) * pow(a, b[i]) % 1337
    }
    return res
}

func pow(x, n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return x % 1337
    }
    return pow(x, n/2) * pow(x, n-n/2) % 1337
}
