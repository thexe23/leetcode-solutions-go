/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var count int
func convertBST(root *TreeNode) *TreeNode {
    count = 0
    helper(root)
    return root
}

func helper(root *TreeNode) {
    if root == nil {
        return
    }
    
    helper(root.Right)
    root.Val += count
    count = root.Val
    helper(root.Left)
}