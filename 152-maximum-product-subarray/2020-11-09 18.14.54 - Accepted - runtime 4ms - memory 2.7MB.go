func maxProduct(nums []int) int {
    curMin, curMax, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        v := nums[i]
        temp := curMax
        curMax = max(v, max(temp * v, curMin * v))
        curMin = min(v, min(temp* v, curMin * v))
        res = max(curMax, res)
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