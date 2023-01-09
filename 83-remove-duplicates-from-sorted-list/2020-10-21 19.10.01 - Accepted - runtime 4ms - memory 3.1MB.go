/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    prev, next := head, head
    for next != nil {
        if next.Val != prev.Val {
            prev.Next = next
            prev = prev.Next
        }
        next = next.Next
    }
    prev.Next = nil
    return head
}