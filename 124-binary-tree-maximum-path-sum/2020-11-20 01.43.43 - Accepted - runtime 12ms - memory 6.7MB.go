/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func maxPathSum(root *TreeNode) int {
    res = -1<< 31
    helper(root)
    return res
}

func helper(root *TreeNode) int{
    if root == nil {
        return 0
    }
    
    left := max(0, helper(root.Left))
    right := max(0, helper(root.Right))
    
    res = max(res, left + right + root.Val)
    return max(left, right) + root.Val
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

