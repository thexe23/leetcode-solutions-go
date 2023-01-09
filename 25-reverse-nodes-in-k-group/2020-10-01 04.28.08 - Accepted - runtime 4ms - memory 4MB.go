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
    hook := head
    count := 0
    nums := []int{}
    for head != nil {
        if count == 0 {
            hook = head
        }
        
        if count <= k - 1 {
            nums = append(nums, head.Val)
        }
        
        if count == k - 1 {
            for i := len(nums) - 1; i >= 0; i--{
                hook.Val = nums[i]
                hook = hook.Next
            }
            nums = []int{}
        }
        count = (count + 1) % k
        head = head.Next
    }
    return dummy.Next
}