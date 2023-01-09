/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    res := Dig(root, 0)
    return res
}

func Dig(root *TreeNode, d int) int {
    if root != nil {
        d++
        d = max(Dig(root.Left, d), Dig(root.Right, d))
    }
    return d
}

func max(x, y int) int {
    if x > y {
        return x
    }
    
    return y
}