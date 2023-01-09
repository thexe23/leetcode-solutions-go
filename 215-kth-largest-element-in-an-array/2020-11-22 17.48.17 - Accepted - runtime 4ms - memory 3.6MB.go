func findKthLargest(nums []int, k int) int {
    return quickSelect(nums, 0, len(nums) - 1, len(nums) - k)
}

func quickSelect(nums []int, start, end, target int) int {
    pivot := randomPatition(nums, start, end)
    if pivot == target {
        return nums[pivot]
    }else if pivot > target {
        return quickSelect(nums, start, pivot - 1, target)
    }
    return quickSelect(nums, pivot + 1, end, target)
}

func randomPatition(nums []int, start, end int) int {
    i := rand.Intn(end - start + 1) + start
    nums[i], nums[end] = nums[end], nums[i]
    return patition(nums, start, end)
}

func patition(nums []int, start, end int) int {
    i := start
    for j := start; j < end; j++ {
        if nums[j] < nums[end] {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    nums[i], nums[end] = nums[end], nums[i]
    return i
}