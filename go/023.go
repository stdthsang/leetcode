func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    for len(lists) > 1 {
        size, temp := len(lists), []*ListNode{}
        for i := 0; i < size-1; i += 2 {
            temp = append(temp, mergeTwoLists(lists[i], lists[i+1]))
        }
        if size%2 != 0 {
            temp = append(temp, lists[size-1])
        }
        lists = temp
    }
    return lists[0]
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
