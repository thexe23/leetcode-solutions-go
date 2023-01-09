func findKthLargest(nums []int, k int) int {
    return quickselect(nums, 0, len(nums) - 1, k)
}

func quickselect(nums []int, lo, hi, k int) int {
    n := len(nums)
    pivot := partition(nums, lo, hi)
    
    if pivot == n - k {
        return nums[pivot]
    }else if pivot > n - k {
        return quickselect(nums, lo, pivot - 1, k)
    }
    
    return quickselect(nums, pivot + 1, hi, k)
}

func partition(nums []int, lo, hi int) int {
    pivot := hi
    i := lo
    for j := lo; j <= hi; j++ {
        if nums[pivot] > nums[j] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[pivot], nums[i] = nums[i], nums[pivot]
    return i
}