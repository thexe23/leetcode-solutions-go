/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        list := []int{}
        n := len(queue)
        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]
            list = append(list, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, list)
    }
    return res
}