func palindromePairs(words []string) [][]int {
    ret, dict := [][]int{}, map[string]int{}
    for i, word := range words {
        dict[reverseString(word)] = i
    }
    if index, ok := dict[""]; ok {
        for i, word := range words {
            if i == index {
                continue
            }
            if isPalindrome(word) {
                ret = append(ret, []int{index, i})
            }
        }
    }
    for i, word := range words {
        wl := len(word)
        for j := 0; j < wl; j++ {
            left, right := word[:j], word[j:]
            if index, ok := dict[left]; ok && isPalindrome(right) && index != i {
                ret = append(ret, []int{i, index})
            }
            if index, ok := dict[right]; ok && isPalindrome(left) && index != i {
                ret = append(ret, []int{index, i})
            }
        }
    }
    return ret
}

func reverseString(str string) string {
    s, size := []byte(str), len(str)
    for i := 0; i < size/2; i++ {
        s[i], s[size-i-1] = s[size-i-1], s[i]
    }
    return string(s)
}

func isPalindrome(str string) bool {
    size := len(str)
    for i := 0; i < size/2; i++ {
        if str[i] != str[size-i-1] {
            return false
        }
    }
    return true
}
