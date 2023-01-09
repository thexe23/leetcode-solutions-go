/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return Symmetric(root, root)
}

func Symmetric(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    
    if (left == nil && right != nil) || (left != nil && right == nil){
        return false
    }
    
    if left.Val != right.Val {
        return false
    }
    return Symmetric(left.Left, right.Right) && Symmetric(left.Right, right.Left)
}