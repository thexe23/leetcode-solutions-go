func fourSum(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums) -3; i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < len(nums) - 2; j++ {
            if j != i + 1 && nums[j] == nums[j - 1]{
                continue
            }
            low := j + 1;
            high := len(nums) - 1
            for low < high {
                if low != j + 1 && nums[low] == nums[low - 1] {
                    low++
                    continue
                }
                if high != len(nums) - 1 && nums[high] == nums[high + 1] {
                    high--
                    continue
                }
                sum := nums[i] + nums[j] + nums[low] + nums[high]
                if sum == target {
                    res = append(res, []int{nums[i], nums[j], nums[low],nums[high]})
                    low++
                }else if sum > target {
                    high--
                }else {
                    low++
                }
            }
        }
    }
    return res
}