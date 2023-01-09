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
        if count % k == 0{
            begin = reverse(begin, head.Next)
            head = begin.Next
            continue
        }
        head = head.Next
    }
    return dummy.Next
}

   /**
     * 0->1->2->3->4->5->6
     * |           |   
     * begin       end
     * after call begin = reverse(begin, end)
     * 
     * 0->3->2->1->4->5->6
     *          |  |
     *      begin end
**/
func reverse(begin, end *ListNode) *ListNode {
    cur := begin.Next
    prev := begin
    first := cur
    for cur != end {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    begin.Next = prev
    first.Next = cur
    return first
}