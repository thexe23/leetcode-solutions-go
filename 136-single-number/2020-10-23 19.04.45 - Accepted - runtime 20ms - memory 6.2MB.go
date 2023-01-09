func singleNumber(nums []int) int {
    m := make(map[int]int)
    for _, n := range nums {
        m[n]++
    }
    for _, n := range nums {
        if m[n] == 1 {
            return n
        }
    }
    return -1
}