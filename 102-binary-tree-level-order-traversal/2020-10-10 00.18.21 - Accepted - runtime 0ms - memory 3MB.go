/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    Traverse(root, &res, 0)
    return res
}

func Traverse(root *TreeNode, res *[][]int, level int){
    if root == nil {
        return
    }
    if len(*res) < level + 1 {
        *res = append(*res, []int{})
    }
    (*res)[level] = append((*res)[level], root.Val)
    Traverse(root.Left, res, level + 1)
    Traverse(root.Right, res, level + 1)
}