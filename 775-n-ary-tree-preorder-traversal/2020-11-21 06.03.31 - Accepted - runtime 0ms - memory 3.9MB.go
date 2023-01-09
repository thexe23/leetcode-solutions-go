/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }
    stack := []*Node{root}
    res := []int{}
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, node.Val)
        for i := len(node.Children) - 1; i >= 0; i-- {
            stack = append(stack, node.Children[i])
        }
    }
    return res
}