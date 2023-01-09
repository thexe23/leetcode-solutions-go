type Solution struct {
    origin []int
}


func Constructor(nums []int) Solution {
    return Solution {
        origin: nums,
    }
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    return this.origin
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    nums := make([]int, len(this.origin))
    copy(nums, this.origin)
    for i := 0; i < len(nums); i++ {
        r := rand.Intn(len(nums))
        nums[i], nums[r] = nums[r], nums[i]
    }
    return nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */