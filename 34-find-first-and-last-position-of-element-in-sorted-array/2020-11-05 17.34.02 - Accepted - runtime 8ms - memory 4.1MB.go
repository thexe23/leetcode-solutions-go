func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    
    fn := func (left bool) int{
        lo, hi := 0, len(nums)
        for lo < hi {
            mid := (hi + lo) / 2
            if (nums[mid] > target) || (left && nums[mid] == target) {
                hi = mid
            }else {
                lo = mid + 1
            }
        }
        return lo
    }
    
    leftIndex := fn(true)
    if leftIndex == len(nums) || nums[leftIndex] != target {
        return res
    }
    res[0] = leftIndex
    res[1] = fn(false) - 1
    return res
    
}