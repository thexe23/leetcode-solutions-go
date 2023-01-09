/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return generate(1, n)
}

func generate(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    res := []*TreeNode{}
    for i := start; i <= end; i++ {
        lefts := generate(start, i - 1)
        rights := generate(i + 1, end)
        for _, l := range lefts {
            for _, r := range rights {
                root := &TreeNode{i, l, r}
                res = append(res, root)
            }
        }
    }
    return res
}