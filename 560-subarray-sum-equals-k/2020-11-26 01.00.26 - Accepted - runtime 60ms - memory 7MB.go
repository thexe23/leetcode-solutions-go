func subarraySum(nums []int, k int) int {
    sum := 0
    res := 0
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum == k {
            res++
        }
        res += m[sum - k]
        m[sum]++
    }
    return res
}