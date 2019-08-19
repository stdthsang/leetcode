func readBinaryWatch(num int) []string {
    var ret []string
    hour, minute := []int{8, 4, 2, 1}, []int{32, 16, 8, 4, 2, 1}
    for i := 0; i <= num; i++ {
        hours := generate(hour, i)
        minutes := generate(minute, num-i)
        for _, h := range hours {
            if h > 11 {
                continue
            }
            for _, m := range minutes {
                if m > 59 {
                    continue
                }
                temp := strconv.Itoa(h)
                if m < 10 {
                    temp += ":0"
                } else {
                    temp += ":"
                }
                temp += strconv.Itoa(m)
                ret = append(ret, temp)
            }
        }
    }
    return ret
}

func generate(num []int, cnt int) []int {
    var ret []int
    helper(num, cnt, 0, 0, &ret)
    return ret
}

func helper(num []int, cnt, pos, out int, ret *[]int) {
    if cnt == 0 {
        *ret = append(*ret, out)
        return
    }
    for i := pos; i < len(num); i++ {
        helper(num, cnt-1, i+1, out+num[i], ret)
    }
}
