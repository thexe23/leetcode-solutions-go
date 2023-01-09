func heightChecker(heights []int) int {
    target := make([]int, len(heights))
    res := 0
    copy(target, heights)
    sort.Ints(target)
    for i, v := range heights {
        if v != target[i] {
            res++
        }
    }
    return res
}