/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    // non-recursive
    if root == nil {
        return nil
    }
    res := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, node.Val)
        root = node.Right
    }
    return res
}