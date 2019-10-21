func minSwapsCouples(row []int) int {
    ret, size := 0, len(row)
    for i := 0; i < size; i += 2 {
        if row[i+1] == (row[i] ^ 1) {
            continue
        }
        ret++
        for j := i + 2; j < size; j++ {
            if row[j] == (row[i] ^ 1) {
                row[i+1], row[j] = row[j], row[i+1]
                break
            }
        }
    }
    return ret
}
