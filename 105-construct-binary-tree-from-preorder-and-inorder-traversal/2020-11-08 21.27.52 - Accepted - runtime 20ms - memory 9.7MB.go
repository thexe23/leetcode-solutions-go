/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    
    root := &TreeNode{Val: preorder[0]}
    
    i := index(inorder, root.Val)
    root.Left = buildTree(preorder[1:i + 1], inorder[:i])
    root.Right = buildTree(preorder[i + 1:], inorder[i + 1:])
    return root
}

func index(nums []int, x int) int {
    for i, v := range nums {
        if x == v {
            return i
        }
    }
    return -1
}