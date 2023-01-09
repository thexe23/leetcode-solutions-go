func trap(height []int) int {
    // DP
    if len(height) == 0 {
        return 0
    }
    res := 0;
    size := len(height)
    left_max := make([]int, size)
    right_max := make([]int, size)
    left_max[0] = height[0]
    for i := 1; i < size; i++ {
        left_max[i] = max(height[i], left_max[i - 1])
    }
    right_max[size - 1] = height[size - 1]
    for j := size - 2; j >=0; j-- {
        right_max[j] = max(height[j], right_max[j + 1])
    }
    
    for k := 0; k < size; k++ {
        res += min(left_max[k], right_max[k]) - height[k]
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    
    return y
}

func min(x, y int) int {
    if x > y {
        return y
    }
    
    return x
}