func findMin(nums []int) int {
    lo, hi := 0, len(nums) - 1
    for lo < hi {
        pivot := lo + (hi - lo) / 2
        if nums[pivot] > nums[hi] {
            lo = pivot + 1
        }else if nums[pivot] < nums[hi] {
            hi = pivot
        }else {
            hi--
        }
    }
    return nums[lo]
}