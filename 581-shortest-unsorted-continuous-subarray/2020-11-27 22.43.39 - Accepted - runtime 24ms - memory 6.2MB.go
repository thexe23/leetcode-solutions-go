func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    left, right, min, max := -1, -2, nums[n - 1], nums[0]
    for i := 1; i < n; i++ {
        if nums[i] >= max {
            max = nums[i]
        }else {
            right = i
        }
        
        if nums[n - 1 - i] <= min {
            min = nums[n - 1 - i]
        }else {
            left = n - 1 - i
        }
    }
    return right - left + 1
}