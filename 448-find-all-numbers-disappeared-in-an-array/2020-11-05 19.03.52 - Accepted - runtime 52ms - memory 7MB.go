func findDisappearedNumbers(nums []int) []int {
    res := []int{}
    for i := 0; i < len(nums); i++ {
        target := abs(nums[i]) - 1
        nums[target] = -abs(nums[target])
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