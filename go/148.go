func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var pre, slow, fast *ListNode
    slow, fast = head, head
    for fast != nil && fast.Next != nil {
        pre, slow, fast = slow, slow.Next, fast.Next.Next
    }
    pre.Next = nil
    return mergeTwoLists(sortList(head), sortList(slow))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    tail := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            temp := l1.Next
            l1.Next = nil
            tail.Next = l1
            l1 = temp
        } else {
            temp := l2.Next
            l2.Next = nil
            tail.Next = l2
            l2 = temp
        }
        tail = tail.Next
    }
    if l1 != nil {
        tail.Next = l1
    }
    if l2 != nil {
        tail.Next = l2
    }
    return dummy.Next
}
