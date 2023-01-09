/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    helper(root, 0)
    return root
}

func helper(root *TreeNode, count int) int {
    if root == nil {
        return count
    }
    root.Val += helper(root.Right, count)
    return helper(root.Left, root.Val)
}