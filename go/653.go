findTarget(root *TreeNode, k int) bool {
    m, stk, cur := map[int]bool{}, []*TreeNode{}, root
    for cur != nil || len(stk) > 0 {
        for cur != nil {
            stk = append(stk, cur)
            cur = cur.Left
        }
        if len(stk) > 0 {
            cur = stk[len(stk)-1]
            if _, ok := m[k-cur.Val]; ok {
                return true
            }
            m[cur.Val] = true
            stk = stk[:len(stk)-1]
            cur = cur.Right
        }
    }
    return false
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

