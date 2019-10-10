func compress(chars []byte) int {
    cur, size := 0, len(chars)
    for i, j := 0, 0; j < size; i = j {
        for j < size && chars[i] == chars[j] {
            j++
        }
        chars[cur] = chars[i]
        cur++
        if j-i == 1 {
            continue
        }
        count := []byte(strconv.Itoa(j - i))
        for _, c := range count {
            chars[cur] = c
            cur++
        }
    }
    return cur
}
