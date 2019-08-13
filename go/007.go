func reverse(x int) int {
    var ret int32
    for x != 0 {
        if abs(ret) > math.MaxInt32 / 10 {
            return 0
        }
        ret = ret * 10 + int32(x) % 10
        x /= 10
    }
    return int(ret)
}

func abs(x int32) int32 {
    if x < 0 {
        return -x
    }
    return x
}
