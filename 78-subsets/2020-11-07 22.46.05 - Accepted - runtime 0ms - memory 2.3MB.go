func subsets(nums []int) [][]int {
    n := len(nums)
    res := [][]int{}
    for mask := 0; mask < (1<<n); mask++ {
        row := []int{}
        for i, v := range nums {
            if (1<<i & mask) != 0 {
                row = append(row, v)
            }
        }
        res = append(res, row)
    }
    return res
}