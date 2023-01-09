func sortColors(nums []int)  {
    red, bule := 0, len(nums) - 1
    for i := 0; i <= bule; i++ {
        if nums[i] == 0 {
            nums[red], nums[i] = nums[i], nums[red]
            red++
        }else if nums[i] == 2 {
            nums[bule], nums[i] = nums[i], nums[bule]
            bule--
            i--
        }
    }
}