/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    dummy := &ListNode{}
    for _, l := range lists {
        res := dummy
        p := dummy.Next
        q := l
        for {
            if p == nil && q == nil {
                break
            }else if p != nil && q == nil {
                res.Next = p
                break
            }else if p == nil && q != nil {
                res.Next = q
                break
            }else if p.Val >= q.Val {
                res.Next = q
                res = res.Next
                q = q.Next
            }else {
                res.Next = p
                res = res.Next
                p = p.Next
            }
        }
}
    return dummy.Next
}