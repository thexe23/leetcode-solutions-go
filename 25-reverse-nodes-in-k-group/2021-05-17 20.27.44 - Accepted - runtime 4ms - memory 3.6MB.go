/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	count := 0
	begin := dummy
	for head != nil {
		count++
		if count%k == 0 {
			begin = reverse(begin, head.Next)
			head = begin.Next
			continue
		}
		head = head.Next
	}
	return dummy.Next
}

func reverse(begin, end *ListNode) *ListNode {
	cur := begin.Next
	prev := begin
	tail := cur
	for cur != end {
		// prev -> cur-> next
		next := cur.Next
		cur.Next = prev // prev <- cur  next
		prev = cur      // old-prev <- prev  next
		cur = next      // old-prev<- prev  cur
	}
	//begin <- xx(tail) <- xx(prev)  end(cur)
	tail.Next = cur
	begin.Next = prev
	// begin -> prev -> tail -> end
	return tail
}
