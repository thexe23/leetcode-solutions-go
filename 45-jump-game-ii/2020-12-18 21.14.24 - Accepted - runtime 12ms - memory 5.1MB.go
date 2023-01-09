func jump(nums []int) int {
    reach, cur := 0, 0
    step := 0
    for i := 0; i < len(nums) - 1; i++ {
        if i + nums[i] > reach {
            reach = i + nums[i]
        }
        if i == cur {
            step++
            cur = reach
        }
    }
    return step
}