func divide(dividend int, divisor int) int {
    //overflow
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    if divisor == 1 {
        return dividend
    }
    if divisor == -1 {
        return -dividend
    }
    sign := 1
    if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
        sign = -1
    }
    
    divAbs, dvAbs, ret := abs(int64(dividend)), abs(int64(divisor)), 0
    for divAbs >= dvAbs {
        temp, shift := dvAbs, 1
        for divAbs >= (temp << 1) {
            temp <<= 1
            shift <<= 1
        }
        divAbs -= temp
        ret += shift

    }
    return ret * sign
}

func abs(a int64) int64 {
    if a < 0 {
        return -a
    }
    return a
}
