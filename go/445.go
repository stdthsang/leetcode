func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var n1, n2 []int
    for ; l1 != nil; l1 = l1.Next {
        n1 = append(n1, l1.Val)
    }
    for ; l2 != nil; l2 = l2.Next {
        n2 = append(n2, l2.Val)
    }
    s1, s2, carry, dummy := len(n1)-1, len(n2)-1, 0, &ListNode{}
    for s1 >= 0 && s2 >= 0 {
        temp := carry + n1[s1] + n2[s2]
        carry = temp / 10
        l := &ListNode{temp % 10, dummy.Next}
        dummy.Next = l
        s1--
        s2--
    }
    for s1 >= 0 {
        temp := carry + n1[s1]
        carry = temp / 10
        l := &ListNode{temp % 10, dummy.Next}
        dummy.Next = l
        s1--
    }
    for s2 >= 0 {
        temp := carry + n2[s2]
        carry = temp / 10
        l := &ListNode{temp % 10, dummy.Next}
        dummy.Next = l
        s2--
    }
    if carry > 0 {
        l := &ListNode{carry, dummy.Next}
        dummy.Next = l
    }
    return dummy.Next
}
