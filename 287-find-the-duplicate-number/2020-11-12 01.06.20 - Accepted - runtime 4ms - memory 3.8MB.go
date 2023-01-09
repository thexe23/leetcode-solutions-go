func findDuplicate(nums []int) int {
    fast, slow := nums[0], nums[0]
    for {
        fast = nums[nums[fast]]
        slow = nums[slow]
        
        if fast == slow {
            break
        }
    }
    head := nums[0]
    for head != slow {
        slow = nums[slow]
        head = nums[head]
    }
    return head
}