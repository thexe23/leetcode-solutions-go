/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    levelOrder(root, 0, &res)
    for i := 1; i < len(res); i = i + 2 {
        res[i] = reverse(res[i])
    }
    return res
}

func levelOrder(root *TreeNode, level int, res *[][]int) {
    if root == nil {
        return
    }
    if len(*res) == level {
        *res = append(*res, []int{})
    }
    (*res)[level] = append((*res)[level], root.Val)

    levelOrder(root.Left, level + 1, res)
    levelOrder(root.Right, level + 1, res)
}

func reverse(nums []int) []int {
    for i := 0; i < len(nums) / 2; i++ {
        nums[len(nums) - 1 - i], nums[i] = nums[i], nums[len(nums) - 1 - i]
    }
    return nums
}