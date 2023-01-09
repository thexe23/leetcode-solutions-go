/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    m := map[*ListNode]int{}
    ptr := head
    for ptr != nil {
        if _, ok := m[ptr]; ok {
            return ptr
        }
        m[ptr]++
        ptr = ptr.Next
    }
    return nil
}