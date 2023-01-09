/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    N := len(lists)
    h := 1
    for h < N {
        for i := 0; i < N - h; i += h * 2 {
            lists[i] = merge(lists[i], lists[i + h])
        }
        h *= 2
    }
    if N > 0 {
        return lists[0]
    }
    return nil
}

func merge(l1, l2 *ListNode) *ListNode {
        dummy := &ListNode{}
        res := dummy
        p := l1
        q := l2
        for {
            if p == nil && q == nil {
                break
            }else if p != nil && q == nil {
                res.Next = p
                break
            }else if p == nil && q != nil {
                res.Next = q
                break
            }else if p.Val >= q.Val {
                res.Next = q
                res = res.Next
                q = q.Next
            }else {
                res.Next = p
                res = res.Next
                p = p.Next
            }
        }
     return dummy.Next
}