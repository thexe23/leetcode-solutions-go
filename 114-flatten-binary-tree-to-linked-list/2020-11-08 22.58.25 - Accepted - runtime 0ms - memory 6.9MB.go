/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode)  {
    var pre *TreeNode = nil
    var helper func(root *TreeNode)
    helper = func(root *TreeNode) {
        if root == nil {
            return
        }
        helper(root.Right)
        helper(root.Left)
        root.Left = nil
        root.Right = pre
        pre = root
    }
    helper(root)
}