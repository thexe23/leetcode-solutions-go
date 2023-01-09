/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    maxDepth(root, &res)
    return res
}

func maxDepth(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    left := maxDepth(root.Left, res)
    right := maxDepth(root.Right, res)
    *res = max(*res, left + right)
    return max(left, right) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}