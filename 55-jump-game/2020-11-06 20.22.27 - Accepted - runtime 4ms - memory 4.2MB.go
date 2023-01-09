func canJump(nums []int) bool {
    reach := 0
    for i := 0; i < len(nums) && i <= reach; i++ {
        reach = max(reach, i + nums[i])
    }
    return reach >= len(nums) - 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
