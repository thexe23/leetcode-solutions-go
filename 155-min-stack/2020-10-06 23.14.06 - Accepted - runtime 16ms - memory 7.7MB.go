type MinStack struct {
    nums []int
    min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        nums: make([]int, 0),
        min: []int{1<<31 -1 },
    }
}


func (this *MinStack) Push(x int)  {
    this.nums = append(this.nums, x)
    if x < this.GetMin() {
        this.min = append(this.min, x)
    }else{
        this.min = append(this.min, this.GetMin())
    }
}


func (this *MinStack) Pop() {
    this.min = this.min[:len(this.min) - 1]
    this.nums = this.nums[:len(this.nums) - 1]
}


func (this *MinStack) Top() int {
    return this.nums[len(this.nums) - 1]
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */