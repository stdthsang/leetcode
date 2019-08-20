func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return head
    }
    cur, dummy := head, &ListNode{Next: head}
    for cnt := 1; cnt < n; cnt++ {
        cur = cur.Next
    }
    pre := dummy
    for cur.Next != nil {
        pre = pre.Next
        cur = cur.Next
    }
    pre.Next = pre.Next.Next
    return dummy.Next
}
