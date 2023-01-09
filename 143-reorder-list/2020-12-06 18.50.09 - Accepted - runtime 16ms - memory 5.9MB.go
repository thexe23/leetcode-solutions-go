/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    fast, slow, prev := head, head, head
    for fast != nil && fast.Next != nil {
        prev = slow
        fast = fast.Next.Next
        slow = slow.Next
    }
    prev.Next = nil
    l1 := head
    l2 := reverse(slow)
    merge(l1, l2)
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

func merge(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    p := dummy
    flag := true
    for l1 != nil && l2 != nil {
        if flag {
            p.Next = l1
            l1 = l1.Next
        }else {
            p.Next = l2
            l2 = l2.Next
        }
        flag = !flag
        p = p.Next
    }
    if l1 == nil {
        p.Next = l2
    }else {
        p.Next = l1
    }
    return dummy.Next
}