func searchInsert(nums []int, target int) int {
    i := 0
    for i < len(nums) && nums[i] <= target {
        i++
    }
    if i == 0 {
        return 0
    }
    if nums[i-1] == target {
        return i - 1
    }else {
        return i
    }
}