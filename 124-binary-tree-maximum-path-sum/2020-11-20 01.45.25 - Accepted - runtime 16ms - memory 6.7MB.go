/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res := -1<< 31
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *int) int{
    if root == nil {
        return 0
    }
    
    left := max(0, helper(root.Left, res))
    right := max(0, helper(root.Right, res))
    
    *res = max(*res, left + right + root.Val)
    return max(left, right) + root.Val
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

