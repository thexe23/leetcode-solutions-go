/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    slow = reverse(slow)
    for slow != nil {
        if slow.Val != head.Val {
            return false
        }
        head = head.Next
        slow = slow.Next
    }
    return true
}

func reverse(head *ListNode) *ListNode{
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}