func subarraySum(nums []int, k int) int {
    sum, res := 0, 0
    m := make(map[int]int)
    m[0] = 1
    for _, v := range nums {
        sum += v
        res += m[sum - k]
        m[sum]++
    }
    return res
}