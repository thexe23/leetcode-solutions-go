/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    res := &TreeNode{}
    if t1 == nil && t2 == nil {
        return nil
    }
    if t1 != nil && t2 == nil {
        res.Val = t1.Val
        res.Left = mergeTrees(t1.Left, nil)
        res.Right = mergeTrees(t1.Right, nil)
    }else if t1 == nil && t2 != nil {
        res.Val = t2.Val
        res.Left = mergeTrees(t2.Left, nil)
        res.Right = mergeTrees(t2.Right, nil)
    }else {
        res.Val = t1.Val + t2.Val
        res.Left = mergeTrees(t1.Left, t2.Left)
        res.Right = mergeTrees(t1.Right, t2.Right)
    }
    return res
}