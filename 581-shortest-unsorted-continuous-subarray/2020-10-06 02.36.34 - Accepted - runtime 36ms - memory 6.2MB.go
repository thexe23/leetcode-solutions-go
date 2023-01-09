func findUnsortedSubarray(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    temp := make([]int, len(nums))
    copy(temp, nums)
    sort.Ints(temp)
    start := len(nums)
    end := 0
    for i := 0; i < len(nums); i++ {
        if temp[i] != nums[i] {
            start = min(start, i)
            end = max(end, i)
        }
    }
    if end > start {
        return end - start + 1
    }
    return 0
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}