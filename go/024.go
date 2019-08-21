func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for head != nil && head.Next != nil {
        temp := head.Next.Next
        head.Next.Next = nil

        tail.Next = head.Next
        tail.Next.Next = head
        head.Next = nil

        tail = tail.Next.Next

        head = temp
    }
    if head != nil {
        tail.Next = head
    }
    return dummy.Next
}
