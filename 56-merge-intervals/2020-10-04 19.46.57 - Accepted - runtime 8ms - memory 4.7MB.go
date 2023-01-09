func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    res := [][]int{}
    
    for i, interval := range intervals {
        if i == 0 {
            res = append(res, interval)
            continue
        }
        
        if interval[0] > res[len(res) - 1][1] {
            res = append(res, interval)
        }
        
        if interval[1] >= res[len(res)- 1][1]{
            res[len(res) - 1] = []int{res[len(res) - 1][0], interval[1]}
        }
    }
    return res
}