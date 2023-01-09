func minSubArrayLen(s int, nums []int) int {
    res := len(nums) + 1
    sum := 0
    for start, end := 0, 0; end < len(nums); end++ {
        sum += nums[end]
        for sum >= s && start <= end{
            res = min(res, end - start + 1)
            sum -= nums[start]
            start++
        }
    }
    
    if res == len(nums) + 1 {
        return 0
    }
    
    return res
}

func min(x, y int) int {
    if x < y {
        return x
    }   
    return y
}