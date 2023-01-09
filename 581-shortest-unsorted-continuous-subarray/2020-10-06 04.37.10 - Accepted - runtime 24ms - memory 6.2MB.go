func findUnsortedSubarray(nums []int) int {
    if len(nums) < 2{
        return 0
    }
    begin, end := -1, -2
    n := len(nums)
    min, max := nums[n - 1], nums[0]
    for i := 1; i < n; i++ {
        if nums[i] >= max {
            max = nums[i]
        }else {
            end = i
        }
        
        if nums[n - 1 - i] <= min {
            min = nums[n - 1 -i]
        }else {
            begin = n - 1 -i
        }
    }
    return end - begin + 1
}