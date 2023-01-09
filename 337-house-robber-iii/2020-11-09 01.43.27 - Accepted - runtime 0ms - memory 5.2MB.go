/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    rob, not := myRob(root)
    return max(rob, not)
}

func myRob(root *TreeNode)(rob, not int) {
    if root == nil {
        return 0, 0
    }
    leftR, leftN := myRob(root.Left)
    rightR, rightN := myRob(root.Right)
    
    return root.Val + leftN + rightN, max(leftN, leftR) + max(rightR, rightN)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}