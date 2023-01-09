/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    return clone(node, map[*Node]*Node{})
}

func clone(node *Node, visited map[*Node]*Node) *Node {
    if node == nil {
        return nil
    }
    
    if v, ok := visited[node]; ok {
        return v
    }
    newNode := &Node{
        Val:node.Val,
        Neighbors: make([]*Node, len(node.Neighbors)),
    }
    visited[node] = newNode
    
    for i, n := range node.Neighbors {
        newNode.Neighbors[i] = clone(n, visited)
    }
    return newNode
}