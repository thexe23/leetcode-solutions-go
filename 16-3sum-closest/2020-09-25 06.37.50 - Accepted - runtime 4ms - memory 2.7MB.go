func threeSumClosest(nums []int, target int) int {
    res := nums[0] + nums[1] + nums[2]
    sort.Ints(nums)
    forloop:
    for i := 0; i < len(nums)-2; i++ {
        low := i + 1
        high := len(nums)-1
        for low < high{
            sum := nums[i] + nums[low] + nums[high]
            if Abs(sum - target) < Abs(res - target) {
                res = sum
            }
            if sum > target {
                high--
            }else if sum == target {
                break forloop
            } else {
                low++
            }
        }
    }
    return res
}
func Abs(x int) int {
    if x > 0{
        return x
    }else{
        return 0 - x
    }
}