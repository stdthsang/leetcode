func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }
    dummy := &ListNode{}
    tail := dummy
    for head != nil {
        kth := findKth(head, k)
        if kth == nil { //length less than k
            tail.Next = head
            break
        }
        next := kth.Next
        kth.Next = nil
        temp := head
        for i := 0; i < k; i++ {
            tmp := head.Next
            head.Next = tail.Next
            tail.Next = head
            head = tmp
        }
        tail = temp
        head = next
    }
    return dummy.Next
}

//return the Kth node, if list length less than k, then return nil
func findKth(head *ListNode, k int) *ListNode {
    for i := 1; i < k && head != nil; i++ {
        head = head.Next
    }
    if head != nil {
        return head
    }
    return nil
}
