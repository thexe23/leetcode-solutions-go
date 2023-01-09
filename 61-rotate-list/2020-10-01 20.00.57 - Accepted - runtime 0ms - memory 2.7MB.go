/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0{
        return head
    }
    nums := []int{}
    res := head
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    head = res
    i := (len(nums) - k % len(nums))%len(nums)
    for ; head != nil; i = (i + 1) % len(nums) {
        head.Val = nums[i]
        head = head.Next
    }
    return res
}