/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    root := helper(nums, 0, len(nums) - 1)
    return root
}

func helper(nums []int, low, high int) *TreeNode {
    if low <= high {
        mid := low + (high - low)/2
        node := &TreeNode{Val: nums[mid]}
        node.Left = helper(nums, low, mid - 1)
        node.Right = helper(nums, mid + 1, high)
        return node
    }
    return nil
}