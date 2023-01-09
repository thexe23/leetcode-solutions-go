func findKthLargest(nums []int, k int) int {
    return QuickSelect(nums, 0, len(nums) - 1, k)
}

func QuickSelect(nums []int, low, high int, k int) int {
    pivot := high
    i := low - 1
    for j := low; j < high; j++ {
        if nums[j] < nums[pivot] {
            i++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    nums[i + 1], nums[pivot] = nums[pivot], nums[i + 1]
    pivot = i + 1
    count := high - pivot + 1
    if count > k {
        return QuickSelect(nums, pivot + 1, high, k)
    }else if count < k {
        return QuickSelect(nums, low, pivot - 1, k - count)
    }
    return nums[pivot]
}