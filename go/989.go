func addToArrayForm(A []int, K int) []int {
    var B []int
    for ; K != 0; K /= 10 {
        B = append(B, K%10)
    }
    for size, i := len(B), 0; i < size/2; i++ {
        B[i], B[size-1-i] = B[size-i-1], B[i]
    }
    carry, s1, s2, ret := 0, len(A)-1, len(B)-1, make([]int, 0)
    for ; s1 >= 0 && s2 >= 0; s1, s2 = s1-1, s2-1 {
        temp := carry + A[s1] + B[s2]
        ret = append(ret, temp%10)
        carry = temp / 10
    }
    for ; s1 >= 0; s1-- {
        temp := carry + A[s1]
        ret = append(ret, temp%10)
        carry = temp / 10
    }
    for ; s2 >= 0; s2-- {
        temp := carry + B[s2]
        ret = append(ret, temp%10)
        carry = temp / 10
    }
    if carry > 0 {
        ret = append(ret, carry)
    }
    for size, i := len(ret), 0; i < size/2; i++ {
        ret[i], ret[size-i-1] = ret[size-i-1], ret[i]
    }
    return ret
}
