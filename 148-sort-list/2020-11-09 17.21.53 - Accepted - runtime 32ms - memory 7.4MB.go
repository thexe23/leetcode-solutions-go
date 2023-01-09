/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    prev, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    l1 := sortList(head)
    l2 := sortList(slow)
    return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    p := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            p.Next = l1
            l1 = l1.Next
        }else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 == nil {
        p.Next = l2
    }else {
        p.Next = l1
    }
    
    return dummy.Next
}