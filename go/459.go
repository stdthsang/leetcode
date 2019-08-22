func repeatedSubstringPattern(s string) bool {
    size := len(s)
    for i := 1; i <= size/2; i++ {
        sub := s[0:i]
        if size%i == 0 {
            temp := ""
            for j := 0; j < size; j += i {
                temp += sub
            }
            if temp == s {
                return true
            }
        }
    }
    return false
}

func repeatedSubstringPattern1(s string) bool {
    size := len(s)
    for i := 1; i <= size/2; i++ {
        sub := s[0:i]
        if size%i == 0 {
            j := i
            for ; j < size; j += i {
                if s[j:i+j] != sub {
                    break
                }
            }
            if j == size && s[j-i:] == sub {
                return true
            }
        }
    }
    return false
}

