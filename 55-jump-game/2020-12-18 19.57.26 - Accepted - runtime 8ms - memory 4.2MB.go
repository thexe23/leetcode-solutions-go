func canJump(nums []int) bool {
    n := len(nums) - 1
    reach := 0
    for i, v := range nums {
        if i <= reach {
            if i + v > reach {
                reach = i + v
            }
        }else {
            return false
        }
        
        if reach >= n {
            return true
        }
    }
    return false
}