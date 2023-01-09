/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    res := head
    nums := []int{}
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    head = res
    n := len(nums) - 1
    for head != nil {
        head.Val = nums[n]
        head = head.Next
        n--
    }
    return res
    
}