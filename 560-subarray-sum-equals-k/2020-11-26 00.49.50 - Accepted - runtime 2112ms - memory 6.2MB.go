func subarraySum(nums []int, k int) int {
    sum := make([]int, len(nums))
    temp := 0
    res := 0
    for i := 0; i < len(nums); i++ {
        temp += nums[i]
        sum[i] = temp
    }
    for i := 0; i < len(sum); i++ {
        if sum[i] == k {
            res++
        }
        for j := 0; j < i; j++ {
            if sum[i] - sum[j] == k {
                res++
            }
        }
    }
    return res
}