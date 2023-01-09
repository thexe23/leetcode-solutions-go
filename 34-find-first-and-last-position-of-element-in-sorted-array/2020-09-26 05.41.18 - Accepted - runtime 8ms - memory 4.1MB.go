func searchRange(nums []int, target int) []int {
    if len(nums) < 1 {
        return []int{-1, -1}
    }
    start := 0
    end := len(nums) - 1
    for start <= end && (nums[start] != target || nums[end] != target){
        if nums[start] < target {
            start++
        }
        if nums[end] > target {
            end--
        }
    }
    if start > end {
        return []int{-1, -1}
    }
    return []int{start, end }
    
}