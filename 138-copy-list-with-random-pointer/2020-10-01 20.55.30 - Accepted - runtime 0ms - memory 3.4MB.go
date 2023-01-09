/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    dummy := &Node{}
    old := head
    p := dummy
    nodeMap := make(map[*Node]*Node)
    for old != nil {
        p.Next = &Node{Val:old.Val}
        nodeMap[old] = p.Next
        old = old.Next
        p = p.Next
    }
    old = head
    p = dummy.Next
    for old != nil {
        if node, ok := nodeMap[old.Random]; ok {
            p.Random = node
        }
        p = p.Next
        old = old.Next
    }
    return dummy.Next
}