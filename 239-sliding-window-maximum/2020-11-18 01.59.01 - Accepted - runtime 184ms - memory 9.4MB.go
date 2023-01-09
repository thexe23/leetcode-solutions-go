func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	queue := []int{}
	for i, v := range nums {
		for len(queue) > 0 && v > queue[len(queue) - 1] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, v)
		if i >= k && nums[i - k] == queue[0] {
			queue = queue[1:]
		}

		if i >= k - 1{
			res = append(res, queue[0])
		}
	}
	return res
}