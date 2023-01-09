func removeElement(nums []int, val int) int {
    sort.Ints(nums)
    i := -1
    for j := 0; j < len(nums); j++ {
        if nums[j] != val {
            i++
            nums[i] = nums[j]
        }
    }
    return i + 1
}