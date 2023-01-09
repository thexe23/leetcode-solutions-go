/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    res := []int{}
    stack := []*TreeNode{}
    for root != nil || len(stack) != 0 {
        for root != nil {
            res = append(res, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        root = node.Right
    }
    return res
}