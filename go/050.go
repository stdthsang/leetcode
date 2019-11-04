func myPow(x float64, n int) float64 {
    res := float64(1)
    for i := n; i != 0; i /= 2 {
        if i&1 == 1 {
            res *= x
        }
        x *= x
    }
    if n < 0 {
        return 1 / res
    }
    return res
}

