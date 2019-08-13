func isPalindrome(head *ListNode) bool {
    v := []int{}
    for head != nil {
        v = append(v, head.Val)
        head = head.Next
    }
    for i, j := 0, len(v)-1; i < j; {
        if v[i] != v[j] {
            return false
        }
        i++
        j--
    }
    return true
}
