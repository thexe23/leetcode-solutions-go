func lengthOfLIS(nums []int) int {
    tails := make([]int, len(nums))
    size := 0
    for _, v := range nums {
        i, j := 0, size
        for i != j {
            m := (i + j) / 2
            if tails[m] < v {
                i = m + 1
            }else {
                j = m
            }
        }
        tails[i] = v
        if i == size {
            size++
        }
    }
    return size
}