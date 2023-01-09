func findDisappearedNumbers(nums []int) []int {
    res := []int{}
    for _, v := range nums {
        nums[abs(v) - 1] = -abs(nums[abs(v) - 1])
    }
    
    for i, v := range nums {
        if v > 0 {
            res = append(res, i + 1)
        }
    }
    
    return res
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}