func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    left, right := make([]int, len(heights)), make([]int, len(heights))
    left[0] = -1
    right[len(heights) - 1] = len(heights)
    res := 0
    
    for i := 1; i < len(heights); i++ {
        p := i - 1
        for p >= 0 && heights[p] >= heights[i] {
            p = left[p]
        }
        left[i] = p
    }
    
    for i := len(heights) - 2; i >= 0; i-- {
        p := i + 1
        for p < len(heights) && heights[p] >= heights[i] {
            p = right[p]
        }
        right[i] = p
    }
    
    for i := 0; i < len(heights); i++ {
        res = max(res, heights[i] * (right[i] - left[i] - 1))
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}