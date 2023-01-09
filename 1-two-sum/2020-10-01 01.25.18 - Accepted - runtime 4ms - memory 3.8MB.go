func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if k, ok := m[target - nums[i]]; ok {
            return []int{i, k}
        }
        m[nums[i]] = i
    }
    return nil
}