/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    return max(myRob(root))
}

func myRob(root *TreeNode)(rob, not int) {
    if root == nil {
        return 0, 0
    }
    leftR, leftN := myRob(root.Left)
    rightR, rightN := myRob(root.Right)
    
    rob = root.Val + leftN + rightN
    not = max(leftN, leftR) + max(rightR, rightN)
    
    return
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}