func twoSum(numbers []int, target int) []int {
	begin, end := 0, len(numbers)-1
	for begin < end {
		sum := numbers[begin] + numbers[end]
		if sum == target {
			return []int{begin + 1, end + 1}
		} else if sum > target {
			end--
			continue
		}
		begin++
	}
	return []int{-1, -1}
}
