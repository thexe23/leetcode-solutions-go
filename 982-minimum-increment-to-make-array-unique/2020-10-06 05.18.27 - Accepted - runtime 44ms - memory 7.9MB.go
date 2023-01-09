func minIncrementForUnique(A []int) int {
    if len(A) < 2 {
        return 0
    }
    count := make([]int, 100000)
    for _, v := range A {
        count[v]++
    }
    res := 0
    for k, v := range count {
        if v > 1 {
            res += count[k] - 1
            count[k + 1] += count[k] - 1
        }
    }
    return res
}