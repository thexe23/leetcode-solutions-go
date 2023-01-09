/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    d1, d2 := &ListNode{}, &ListNode{}
    l1, l2 := d1, d2
    for head != nil {
        if head.Val < x {
            l1.Next = &ListNode{Val:head.Val}
            l1 = l1.Next
        }else {
            l2.Next = &ListNode{Val:head.Val}
            l2 = l2.Next
        }
        head = head.Next
    }
    l1.Next = d2.Next
    return d1.Next
}