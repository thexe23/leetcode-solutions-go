func maxProduct(nums []int) int {
    curMin, curMax, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        temp := curMax
        curMax = max(nums[i], max(temp * nums[i], curMin * nums[i]))
        curMin = min(nums[i], min(temp* nums[i], curMin * nums[i]))
        if curMax > res {
            res = curMax
        }
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}