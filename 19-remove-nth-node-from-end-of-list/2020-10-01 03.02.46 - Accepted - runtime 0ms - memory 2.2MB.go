/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    res, p := head, head
    sum := 0
    for p != nil {
        p = p.Next
        sum++
    }
    
    if sum == n {
        return head.Next
    }
    for i := 1; i < sum - n; i++ {
        head = head.Next
    }
    if head.Next != nil {
        temp := head.Next.Next
        head.Next = temp
    }else {
        head = nil
    }
    
    return res
}