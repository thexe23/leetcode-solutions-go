func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    res[0] = 1
    for i := 1; i < len(nums); i++ {
        res[i] = res[i - 1] * nums[i - 1]
    }
    right := 1
    for j := len(nums) - 1; j >= 0; j-- {
        res[j] *= right
        right *= nums[j]
    }
    return res
}