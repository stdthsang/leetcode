func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy, carry := &ListNode{}, 0
    tail := dummy
    for l1 != nil && l2 != nil {
        temp := l1.Val + l2.Val + carry
        tail.Next = &ListNode{temp % 10, nil}
        carry = temp / 10
        l1, l2, carry, tail = l1.Next, l2.Next, temp/10, tail.Next
    }
    for l1 != nil {
        temp := l1.Val + carry
        tail.Next = &ListNode{temp % 10, nil}
        carry = temp / 10
        l1, carry, tail = l1.Next, temp/10, tail.Next
    }
    for l2 != nil {
        temp := l2.Val + carry
        tail.Next = &ListNode{temp % 10, nil}
        carry = temp / 10
        l2, carry, tail = l2.Next, temp/10, tail.Next
    }
    if carry > 0 {
        tail.Next = &ListNode{carry, nil}
    }
    return dummy.Next
}
