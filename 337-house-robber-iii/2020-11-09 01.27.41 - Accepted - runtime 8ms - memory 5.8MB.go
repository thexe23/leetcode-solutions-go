/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    res := myRob(root)
    return max(res[0], res[1])
}

func myRob(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }
    
    left := myRob(root.Left)
    right := myRob(root.Right)
    
    rob := root.Val + left[1] + right[1]
    notRob := max(left[0], left[1]) + max(right[0], right[1])
    
    return []int{rob, notRob}
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}