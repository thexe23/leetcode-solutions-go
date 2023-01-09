func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    
    return max(myRob(nums[:n - 1]), myRob(nums[1: n]))
}

func myRob(nums []int) int{
    prev, cur := 0, 0
    for _, v := range nums {
        temp := cur
        cur = max(cur, prev + v)
        prev = temp
    }
    return cur
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
