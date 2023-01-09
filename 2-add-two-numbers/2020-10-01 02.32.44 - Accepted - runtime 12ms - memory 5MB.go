/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    re := 0
    res := &ListNode{}
    head := res
    for l1 != nil || l2 != nil || re != 0 {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        sum += re        
        res.Next = &ListNode{sum % 10, nil}
        re = sum / 10
        
        res = res.Next
    }
    return head.Next
}