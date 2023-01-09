func nextPermutation(nums []int)  {
    if len(nums) < 2 {
        return
    }
    i := len(nums) - 2;
    for i >= 0 && nums[i] >= nums[i + 1]{
        i--
    }
    if i >= 0 {
        for j := len(nums) - 1; j > i; j--{
            if nums[j] > nums[i] {
                temp := nums[j]
                nums[j] = nums[i]
                nums[i] = temp
                break
            }
        }
    }
    head := i + 1
    tail := len(nums) - 1
    for head < tail {
        temp := nums[head]
        nums[head] = nums[tail]
        nums[tail] = temp
        head++
        tail--
    }
}