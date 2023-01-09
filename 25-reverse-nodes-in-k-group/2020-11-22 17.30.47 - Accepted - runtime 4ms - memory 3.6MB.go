/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next:head}
    start, end := dummy, head
    count := 0
    for end != nil {
        count++
        end = end.Next
        if count % k == 0 {
            start = reverse(start, end)
            end = start.Next
        }
    }
    return dummy.Next
}

func reverse(start, end *ListNode) *ListNode {
    curr := start.Next
    prev := start
    tail := curr
    
    for curr != end {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
// dummy--->1--->2---->3
// dummy<---1<---2     3
    tail.Next = end
    start.Next = prev
    return tail
}