func findKthLargest(nums []int, k int) int {
    n := len(nums)
    return quickSelect(nums, 0, n - 1, n - k)
}

func quickSelect(nums []int, low, high, target int) int {
    pivot := partition(nums, low, high)
    if pivot == target {
        return nums[pivot]
    }else if pivot < target {
        return quickSelect(nums, pivot + 1, high, target)
    }
    return quickSelect(nums, low, pivot - 1, target)
}

func partition(nums []int, low, high int) int {
    // pick high as pivot
    i := low
    for j := low; j < high; j++ {
        if nums[j] < nums[high] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[high], nums[i] = nums[i], nums[high]
    return i
}