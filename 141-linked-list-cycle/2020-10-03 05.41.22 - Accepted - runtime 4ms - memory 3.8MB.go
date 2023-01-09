/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    
    fast, slow := head.Next, head
    for {
        if fast == nil || fast.Next == nil {
            return false
        }
        
        if slow == fast {
            break
        }
        
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}