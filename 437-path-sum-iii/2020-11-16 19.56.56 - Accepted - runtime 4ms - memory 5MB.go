/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    pathSum := make(map[int]int)
    pathSum[0]++
    res := 0
    visit(root, sum, 0, &res, pathSum)
    return res
}

func visit(root *TreeNode, sum int, cur int, res *int, pathSum map[int]int) {
    if root == nil {
        return 
    }
    
    cur += root.Val
    *res += pathSum[cur - sum]
    pathSum[cur]++
    visit(root.Left, sum, cur, res, pathSum)
    visit(root.Right, sum, cur, res, pathSum)
    pathSum[cur]--
}