/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    res := true
    maxDepth(root, &res)
    return res
}

func maxDepth(root *TreeNode, res *bool) int {
    if root == nil {
        return 0
    }
    
    l := maxDepth(root.Left, res)
    r := maxDepth(root.Right, res)
    
    if abs(l - r) > 1 {
        *res = false
    }
    
    return 1 + max(l, r)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}