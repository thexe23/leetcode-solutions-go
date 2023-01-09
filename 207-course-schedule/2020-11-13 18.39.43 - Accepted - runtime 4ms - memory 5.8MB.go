func canFinish(numCourses int, prerequisites [][]int) bool {
    edges := make([][]int, numCourses)
    inedge := make([]int, numCourses)
    res := []int{}
    q := []int{}
    
    for _, v := range prerequisites {
        edges[v[1]] = append(edges[v[1]], v[0])
        inedge[v[0]]++
    }
    
    for i, v := range inedge {
        if v == 0 {
            q = append(q, i)
        }
    }
    
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        res = append(res, u)
        for _, v := range edges[u] {
            inedge[v]--
            if inedge[v] == 0 {
                q = append(q, v)
            }
        }
    }
    return len(res) == numCourses
}