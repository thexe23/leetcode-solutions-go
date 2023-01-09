/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    fast, slow := head, head
    length := 0
    for slow != nil {
        length++
        slow = slow.Next
    }
    slow = head
    k %= length
    if k == 0 {
        return head
    }
    for i := 0; i < k; i++ {
        fast = fast.Next
    }
    
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    res := slow.Next
    fast.Next = head
    slow.Next = nil
    return res
}