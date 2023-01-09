func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    beg := -1
    end := -2 
    min := nums[n-1]
    max := nums[0]
    for i := 1; i < n; i++ {
      max = maxInt(max, nums[i])
      min = minInt(min, nums[n-1-i])
        if (nums[i] < max) {
            end = i
        }
        if nums[n-1-i] > min {
            beg = n-1-i
        }
    }
    return end - beg + 1;
}

func maxInt(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func minInt(x, y int) int {
    if x > y {
        return y
    }
    return x
}
