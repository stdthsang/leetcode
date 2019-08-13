func isPalindrome(x int) bool {
    if x < 0 || x%10 == 0 && x != 0 {
        return false
    }
    y := 0
    for x > y {
        y = y*10 + x%10
        x /= 10
    }
    return x == y || x == y/10
}

func isPalindrome2(x int) bool {
    if x < 0 || x%10 == 0 && x != 0 {
        return false
    }
    x1, y := x, 0
    for x1 != 0 {
        if y > math.MaxInt32/10 {
            return false
        }
        y = y*10 + x1%10
        x1 /= 10
    }
    return y == x
}
