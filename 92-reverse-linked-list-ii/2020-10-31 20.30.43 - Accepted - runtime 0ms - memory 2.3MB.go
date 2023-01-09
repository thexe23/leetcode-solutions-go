/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy
    count := 1
    for count < m {
        prev = prev.Next
        count++
    }
    cur := prev.Next
    next := cur.Next
    for count < n {
        cur.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
        next = cur.Next
        count++
    }
    return dummy.Next
}