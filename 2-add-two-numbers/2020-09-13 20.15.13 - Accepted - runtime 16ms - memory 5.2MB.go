/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    carry := 0
    curr := dummy
    for(l1 != nil || l2 != nil || carry != 0){
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        
        curr.Next = new(ListNode)
        curr.Next.Val = carry % 10
        curr = curr.Next
        carry /= 10
    }
    return dummy.Next
}