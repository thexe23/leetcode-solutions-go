/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    levels := [][]int{}
    res := []float64{}
    levelTravese(root, 0, &levels)
    for _, l := range levels {
        sum := 0
        n := len(l)
        for _, n := range l {
            sum += n
        }
        res = append(res, float64(sum) / float64(n))
    }
    return res
}

func levelTravese(root *TreeNode, level int, res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) < level + 1{
        *res = append(*res, []int{})
    }
    
    (*res)[level] = append((*res)[level], root.Val)
    
    levelTravese(root.Left, level + 1, res)
    levelTravese(root.Right, level + 1, res)
}