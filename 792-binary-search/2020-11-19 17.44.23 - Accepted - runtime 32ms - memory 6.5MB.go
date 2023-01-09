func search(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        if target == nums[mid] {
            return mid
        }else if target > nums[mid] {
            lo = mid + 1
        }else {
            hi = mid - 1
        }
    }
    return -1
}