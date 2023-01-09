func threeSumClosest(nums []int, target int) int {
    res := nums[0] + nums[1] + nums[2]
    for first := 0; first < len(nums)-2; first ++ {
        for second := first + 1; second < len(nums) - 1; second++ {
            for third := second + 1; third < len(nums); third++ {
                sum := nums[first] + nums[second] + nums[third]
                if Abs(sum - target) < Abs(res - target){
                    res = sum
                }
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