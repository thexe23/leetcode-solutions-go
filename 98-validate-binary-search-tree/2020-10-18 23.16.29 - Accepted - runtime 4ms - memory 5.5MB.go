/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValid(root,nil, nil)
}

func isValid(root *TreeNode, low, high interface{}) bool {
    if root == nil {
        return true
    }
    
    if low != nil && root.Val <= low.(int) {
        return false
    }
    
    if high != nil && root.Val >= high.(int) {
        return false
    }
    
    return isValid(root.Left, low, root.Val) && isValid(root.Right, root.Val, high)
}