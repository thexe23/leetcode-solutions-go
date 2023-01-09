/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var count int
    stack := []*TreeNode{}
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        count++
        if count == k {
            return node.Val
        }
        root = node.Right
    }
    return -1
}