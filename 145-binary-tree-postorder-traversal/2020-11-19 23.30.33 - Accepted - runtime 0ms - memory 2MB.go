/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    //iterative
    if root == nil {
        return nil
    }
    res := []int{}
    stack := []*TreeNode{}
    var prev *TreeNode
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        node := stack[len(stack) - 1]
        if node.Right == nil || node.Right == prev {
            stack = stack[:len(stack) - 1]
            res = append(res, node.Val)
            prev = node
        }else {
            root = node.Right
        }
    }
    return res
}