func isAnagram(s string, t string) bool {
    size1, size2 := len(s), len(t)
    if size1 != size2 {
        return false
    }
    dict := make([]byte, 26)
    for _, c := range s {
        dict[c-'a']++
    }
    for _, c := range t {
        dict[c-'a']--
    }
    for _, c := range dict {
        if c != 0 {
            return false
        }
    }
    return true
}

