/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    res := dummy
    p := l1
    q := l2
    for p != nil && q != nil {
        if p.Val >= q.Val {
            res.Next = q
            q = q.Next
            res = res.Next
        }else {
            res.Next = p
            p = p.Next
            res = res.Next
        }
    }
    
    if p == nil{
        res.Next = q
    }
    
    if q == nil {
        res.Next = p
    }
    
    return dummy.Next
}