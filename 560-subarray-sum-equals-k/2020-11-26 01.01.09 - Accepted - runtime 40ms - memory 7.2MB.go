func subarraySum(nums []int, k int) int {
    sum := 0
    res := 0
    m := make(map[int]int)
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        res += m[sum - k]
        m[sum]++
    }
    return res
}