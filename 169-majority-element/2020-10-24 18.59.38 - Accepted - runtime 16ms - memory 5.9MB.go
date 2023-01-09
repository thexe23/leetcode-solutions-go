func majorityElement(nums []int) int {
    var candidate int
    count := 0
    for i := 0; i < len(nums); i++ {
        if count == 0 {
            candidate = nums[i]
        }
        if nums[i] == candidate {
            count++
        }else {
            count--
        }
    }
    return candidate
}