/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    res := [][]int{}
    helper(&res, root, 0)
    res = reverse(res)
    return res
}

func helper(res *[][]int, root *TreeNode, level int) {
    if root == nil {
        return
    }
    
    if level >= len(*res) {
        *res = append(*res, []int{})
    }
    
    (*res)[level] = append((*res)[level], root.Val)
    helper(res, root.Left, level + 1)
    helper(res, root.Right, level + 1)
    return
}

func reverse(mat [][]int) [][]int {
    res := make([][]int, len(mat))
    for i := len(mat) - 1; i >=0; i-- {
        res[len(mat) - 1 - i] =  mat[i]
    }
    return res
}