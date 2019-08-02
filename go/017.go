func twoSum(numbers []int, target int) []int {
    size, ret := len(numbers), []int{}
    for i, j := 0, size - 1; i < j; {
        if numbers[i] + numbers[j] == target {
            ret = append(ret, i + 1, j + 1)
            return ret
        } else if numbers[i] + numbers[j] < target {
            i++
        } else {
            j--
        }
    }   
    return ret
}
