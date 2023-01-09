/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    return Inorder(root, res)
}

func Inorder(root *TreeNode, res []int) []int{
    if root != nil {
        if root.Left != nil {
            res = Inorder(root.Left, res)
        }
        
        res = append(res, root.Val)
        
        if root.Right != nil {
            res = Inorder(root.Right, res)
        }
    }
    return res
}