func threeSum(nums []int) [][]int {
    if len(nums) < 2 {
        return [][]int{}
    }
    
    res := [][]int{}
    
    sort.Ints(nums)
    
    for i := 0; i < len(nums)-2; i++{
        target, left, right := 0 - nums[i], i + 1, len(nums) - 1
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            }else if sum > target {
                right--
            }else if sum < target {
                left++
            }
        }
    }
    return res
}