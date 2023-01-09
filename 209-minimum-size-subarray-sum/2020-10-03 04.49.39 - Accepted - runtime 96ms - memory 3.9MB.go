func minSubArrayLen(s int, nums []int) int {
    start, end := 0, 0
    res := len(nums) + 1
    for end < len(nums) {
        if sum(nums[start:end + 1]) >= s {
            res = min(res, end - start + 1)
            start++
        }else{
            end++
        }
    }
    if res == len(nums) + 1 {
        return 0
    }
    return res
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func sum(n []int) int {
    sum := 0
    for _, v := range n {
        sum += v
    }
    return sum
}