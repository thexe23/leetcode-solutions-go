/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    nodes := []*TreeNode{root}
    dep := 1
    for {
        newNodes := []*TreeNode{}
        
        for _, node := range nodes {
            if node.Left == nil && node.Right == nil {
                return dep
            }
            
            if node.Right != nil {
                newNodes = append(newNodes, node.Right)
            }
            
            if node.Left != nil {
                newNodes = append(newNodes, node.Left)
            }
        }
        dep++
        nodes = newNodes
    }
    return dep
}
